package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"user/config"
	"user/core"
	srv "user/services"
)

func main() {
	config.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("47.107.39.175:2379"),
		)
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(":8082"),
		micro.Registry(etcdReg),
		)
	microService.Init()

	_  = srv.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	_ = microService.Run()
}