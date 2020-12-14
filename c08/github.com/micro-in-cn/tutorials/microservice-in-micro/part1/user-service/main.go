package main

import (
	"fmt"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/model"
	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.service.user"),
		micro.Registry(micReg),
		micro.Version("lastest"),
	)
	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			model.Init()
			handler.Init()
			return nil
		}),
	)
	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
