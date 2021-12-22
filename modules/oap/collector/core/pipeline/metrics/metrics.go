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

package metrics

import (
	"context"
	"fmt"
	"sync"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda/modules/oap/collector/core/model"
)

type Pipeline struct {
	mReceivers  []model.MetricReceiver
	mProcessors []model.MetricProcessor
	mExporters  []model.MetricExporter

	Log logs.Logger
}

func NewPipeline(log logs.Logger) *Pipeline {
	return &Pipeline{Log: log}
}

func (p *Pipeline) InitComponents(receivers, processors, exporters []model.Component) error {
	rs, err := p.rsFromComponent(receivers)
	if err != nil {
		return err
	}
	prs, err := p.prsFromComponent(processors)
	if err != nil {
		return err
	}
	es, err := p.esFromComponent(exporters)
	if err != nil {
		return err
	}

	p.mReceivers = rs
	p.mProcessors = prs
	p.mExporters = es
	return nil
}

func (p *Pipeline) rsFromComponent(coms []model.Component) ([]model.MetricReceiver, error) {
	res := make([]model.MetricReceiver, 0, len(coms))
	for _, com := range coms {
		r, ok := com.(model.MetricReceiver)
		if !ok {
			return nil, fmt.Errorf("invalid component<%s> type", com.ComponentID())
		}
		res = append(res, r)
	}
	return res, nil
}
func (p *Pipeline) prsFromComponent(coms []model.Component) ([]model.MetricProcessor, error) {
	res := make([]model.MetricProcessor, 0, len(coms))
	for _, com := range coms {
		r, ok := com.(model.MetricProcessor)
		if !ok {
			return nil, fmt.Errorf("invalid component<%s> type", com.ComponentID())
		}
		res = append(res, r)
	}
	return res, nil
}
func (p *Pipeline) esFromComponent(coms []model.Component) ([]model.MetricExporter, error) {
	res := make([]model.MetricExporter, 0, len(coms))
	for _, com := range coms {
		r, ok := com.(model.MetricExporter)
		if !ok {
			return nil, fmt.Errorf("invalid component<%s> type", com.ComponentID())
		}
		res = append(res, r)
	}
	return res, nil
}

func (p *Pipeline) StartStream(ctx context.Context) {
	out := make(chan model.Metrics)
	in := make(chan model.Metrics)
	go p.StartExporters(ctx, out)

	go p.startProcessors(ctx, in, out)

	p.startReceivers(ctx, in)
}

func (p *Pipeline) StartExporters(ctx context.Context, out <-chan model.Metrics) {
	for {
		select {
		case data := <-out:
			var wg sync.WaitGroup
			wg.Add(len(p.mExporters))
			for _, e := range p.mExporters {
				go func(exp model.MetricExporter, ms model.Metrics) {
					defer wg.Done()
					err := exp.ExportMetrics(ms)
					if err != nil {
						p.Log.Errorf("Exporter<%s> export data error: %s", exp.ComponentID(), err)
					}
				}(e, data.Clone())
			}
			wg.Wait()
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pipeline) startProcessors(ctx context.Context, in <-chan model.Metrics, out chan<- model.Metrics) {
	for {
		select {
		case data := <-in:
			for _, pr := range p.mProcessors {
				tmp, err := pr.ProcessMetrics(data)
				if err != nil {
					p.Log.Errorf("Processor<%s> process data error: %s", pr.ComponentID(), err)
					continue
				}
				data = tmp
			}
			// wait forever
			select {
			case out <- data:
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pipeline) startReceivers(ctx context.Context, in chan<- model.Metrics) {
	for _, r := range p.mReceivers {
		consumer := func(ms model.Metrics) {
			select {
			case in <- ms:
			case <-ctx.Done():
			}
		}
		r.RegisterConsumeFunc(consumer)
	}
}
