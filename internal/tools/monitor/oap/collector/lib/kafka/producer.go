// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafka

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Shopify/sarama"
)

var (
	mu             sync.Mutex
	sharedProducer *AsyncProducer // producer is thread safe, so we shared
)

// ProducerConfig try to keep same semantics as kafka official doc
type producerConfig struct {
	MaxBlockMs        time.Duration `file:"max.block.ms" default:"30s"`
	LingerMs          time.Duration `file:"linger.ms" default:"50ms"`
	RequestTimeoutMs  time.Duration `file:"request.timeout.ms" default:"30s"`
	BatchSize         int           `file:"batch.size" default:"16384"`
	BufferMemory      int           `file:"buffer.memory" default:"33554432"`
	ChannelBufferSize int           `file:"channel_buffer_size" default:"500" env:"KAFKA_P_CHANNEL_BUFFER_SIZE"`
}

type ProducerOption interface{}

// NewProducer default shared
func (p *provider) NewProducer() (*AsyncProducer, error) {
	mu.Lock()
	defer mu.Unlock()
	if sharedProducer != nil {
		atomic.AddUint32(sharedProducer.ref, 1)
		return sharedProducer, nil
	}

	cfg := sarama.NewConfig()
	cfg.Version = p.protoVersion
	cfg.ClientID = p.Cfg.ClientID
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.ChannelBufferSize = p.Cfg.Producer.ChannelBufferSize
	cfg.Producer.Flush.MaxMessages = p.Cfg.Producer.BatchSize
	cfg.Producer.Flush.Frequency = p.Cfg.Producer.LingerMs
	cfg.Producer.Flush.Bytes = p.Cfg.Producer.BufferMemory

	producer, err := sarama.NewAsyncProducer(p.Brokers(), cfg)
	if err != nil {
		return nil, fmt.Errorf("create async producer: %w", err)
	}

	ref := uint32(1)
	sharedProducer = &AsyncProducer{
		ref:          &ref,
		producer:     producer,
		blockTimeout: p.Cfg.Producer.MaxBlockMs,
	}

	go func() {
		for pe := range producer.Errors() {
			if pe.Err == nil {
				continue
			}
			p.Log.Sub("producer").Errorf("failed to produce, topic: %s, err: %s", pe.Msg.Topic, pe.Err)
		}
	}()
	return sharedProducer, nil
}

type AsyncProducer struct {
	ref          *uint32
	blockTimeout time.Duration
	producer     sarama.AsyncProducer
}

func (a *AsyncProducer) Write(data interface{}) error {
	switch msg := data.(type) {
	case *sarama.ProducerMessage:
		return a.send(msg)
	default:
		return fmt.Errorf("unsupported data type: %T", data)
	}
}

func (a *AsyncProducer) send(pmsg *sarama.ProducerMessage) error {
	select {
	case a.producer.Input() <- pmsg:
	case <-time.After(a.blockTimeout):
		return fmt.Errorf("produce message block timout")
	}
	return nil
}

func (a *AsyncProducer) WriteN(data ...interface{}) (int, error) {
	offset := 0
	for ; offset < len(data); offset++ {
		err := a.Write(data[offset])
		if err != nil {
			return offset, err
		}
	}
	return offset, nil
}

func (a *AsyncProducer) Close() error {
	atomic.AddUint32(a.ref, ^uint32(0))
	if atomic.LoadUint32(a.ref) > 0 {
		return nil
	}
	err := a.producer.Close()
	if err != nil {
		return fmt.Errorf("close produer: %w", err)
	}
	return nil
}
