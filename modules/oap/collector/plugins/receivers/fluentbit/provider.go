package fluentbit

import (
	"fmt"
	"net/http"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/httpserver"
	"github.com/erda-project/erda/modules/oap/collector/common"
	"github.com/erda-project/erda/modules/oap/collector/core/model"
	"github.com/erda-project/erda/modules/oap/collector/plugins"
	"github.com/labstack/echo"
)

var providerName = plugins.WithPrefixReceiver("fluent-bit")

type config struct {
	FLBKeyMappings flbKeyMappings `file:"flb_key_mappings"`
}

// This is the key mappings config of erda's internal model to fluent-bit's output format
type flbKeyMappings struct {
	TimeUnixNano string `file:"time_unix_nano" default:"time"`
	Name         string `file:"name" default:"source"`
	Content      string `file:"log" default:"log"`
	Severity     string `file:"severity" default:"level"`
	// parse kubernetes's metadata from `kubernetes` field
	Kubernetes string `file:"kubernetes" default:"kubernetes"`
}

// +provider
type provider struct {
	Cfg    *config
	Log    logs.Logger
	Router httpserver.Router `autowired:"http-router"`

	consumerFunc model.ObservableDataConsumerFunc
}

func (p *provider) ComponentID() model.ComponentID {
	return model.ComponentID(providerName)
}

func (p *provider) RegisterConsumer(consumer model.ObservableDataConsumerFunc) {
	p.consumerFunc = consumer
	return
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	p.Router.POST("/api/v1/collect/fluent-bit", p.flbHandler)
	return nil
}

func (p *provider) flbHandler(ctx echo.Context) error {
	if p.consumerFunc == nil {
		return ctx.NoContent(http.StatusOK)
	}
	req := ctx.Request()
	buf, err := common.ReadBody(req)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("read body err: %s", err))
	}

	data, err := p.convertToLogs(buf)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("convertToLogs err: %s", err))
	}
	p.consumerFunc(data)

	return ctx.NoContent(http.StatusOK)
}

func init() {
	servicehub.Register(providerName, &servicehub.Spec{
		Services: []string{
			providerName,
		},
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
