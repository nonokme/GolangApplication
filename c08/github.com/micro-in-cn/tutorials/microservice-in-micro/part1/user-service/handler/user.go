package handler

import (
	"context"
	us "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/model/user"
	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/proto/user"
	log "github.com/micro/go-micro/v2/logger"
)

type Service struct{}

var (
	userService us.Service
)

func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}
	rsp.User = user
	rsp.Success = true
	return nil
}

// 以下为自动生成的代码
//type User struct{}
//// Call is a single request handler called via client.Call or the generated client code
//func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
//	log.Info("Received User.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}
//
//// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
//	log.Infof("Received User.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Infof("Responding: %d", i)
//		if err := stream.Send(&user.StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Infof("Got ping %v", req.Stroke)
//		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}
