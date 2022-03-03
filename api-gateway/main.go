package main

import (
	services "api-gateway/services"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"),
		)

	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
		)


	userService := services.NewUserService("rpcUserService", userMicroService.Client())

	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
		)
	taskService := services.NewTaskService("rpcTaskService", taskMicroService.Client())
	server := web.NewService(
		web.Address("127.0.0.1:4000"),
		web.Name("httpService"),
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	_ = server.Init()
	_ = server.Run()
}