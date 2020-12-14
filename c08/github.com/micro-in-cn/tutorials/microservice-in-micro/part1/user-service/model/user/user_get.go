package user

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/basic/db"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/proto/user"
	"github.com/prometheus/common/log"
)

func(s *service) QueryUserByName(userName string)(ret *proto.User,err error){
	queryString:="select user_id,user_name,pwd from user where user_name=?"
	o:=db.GetDB()
	ret=&proto.User{}

	err=o.QueryRow(queryString,userName).Scan(&ret.Id,&ret.Name,&ret.Pwd)
	if err!=nil{
		log.Errorf("[QueryUserByName] 查询数据失败,err:%s",err)
		return
	}
	return
}
