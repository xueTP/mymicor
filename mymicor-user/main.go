package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"log"
	"mymicor/mymicor-user/data"
	userserver "mymicor/mymicor-user/server"
)

var conn *gorm.DB
var pubSub broker.Broker
var publish micro.Publisher

func main() {
	// gorm create conn
	var err error
	conn, err = data.CreateConnection()
	if err != nil {
		logrus.Errorf("create gorm connection error: %v", err)
	}
	defer conn.Close()
	// server init
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.WrapHandler(AuthHandel),
	)
	srv.Init()

	// go-micor broker 貌似默认使用nats 消息中间件，但是
	// 我本地跑买有开启nats :4222 服务，但是还是能够接收
	// 到发布事件
	pubSub = srv.Server().Options().Broker
	//go-micro 提供的pushlisher 服务 真正棒的部分在于它利用了 protobuf 的定义
	publish = micro.NewPublisher("user.create", srv.Client())

	userServer := userserver.NewUserServer(conn, pubSub, publish)
	// Register handler
	pd.RegisterUserServiceHandler(srv.Server(), userServer)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthHandel(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		if req.Method() == "UserService.Auth" || req.Method() == "UserService.Create" {
			return fn(ctx, req, rsp)
		}
		data, ok := metadata.FromContext(ctx)
		if ! ok {
			return errors.New("no auth meta-data found in request")
		}
		token := data["Token"]
		log.Println("token is ", token)
		var res *pd.Token
		err := userserver.NewUserServer(conn, pubSub, publish).ValidateToken(context.TODO(), &pd.Token{Token: token}, res)
		if err != nil {
			return err
		}
		return fn(ctx, req, rsp)
	}
}
