package user

import (
	"fmt"
	"sync"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/proto/user"
)

var (
	s *service
	m sync.RWMutex
)

// service服务
type service struct{

}

type Service interface {
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// 暴露服务
func GetService() (*service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] Getservice 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
