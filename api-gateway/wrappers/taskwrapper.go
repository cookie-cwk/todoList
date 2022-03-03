package wrappers

import (
	service "api-gateway/services"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4/client"
	"strconv"
)

func NewTask(id uint64, name string) *service.TaskModel {
	return &service.TaskModel {
		Id:         id,
		Title:      name,
		Content:    "响应超时",
		StartTime:  1000,
		EndTime:    1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

func DefaultTasks(resp interface{}) {
	models := make([]*service.TaskModel, 0)
	for i := 0; i < 10; i++ {
		models = append(models, NewTask(uint64(i), "降级备忘录" + strconv.Itoa(20 + int((i)))))
	}
	result := resp.(*service.TaskListResponse)
	result.TaskList = models
}

type TaskWrapper struct {
	client.Client
}

func (wrapper *TaskWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  50,
		SleepWindow:            5000,
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		DefaultTasks(resp)
		return err
	})
}

func NewTaskWrapper(c client.Client) client.Client{
	return &TaskWrapper{c}
}