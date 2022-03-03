package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4/client"
)

type userWrapper struct {
	client.Client
}


func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 20,
		SleepWindow:            5000,
		ErrorPercentThreshold:  50,
	}
	hystrix.ConfigureCommand(cmdName, config)
	return  hystrix.Do(cmdName, func() error {
		return 	wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

func NewUserWrapper(c client.Client) client.Client {
	return &userWrapper{c}
}