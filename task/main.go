package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"task/conf"
	"task/core"
	services "task/services"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)

	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
		)
	microService.Init()
	_ = services.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	_ = microService.Run()
}
