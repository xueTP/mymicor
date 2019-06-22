package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"log"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)
	srv.Init()

	// 监听发布用户创建事件 broker 监听
	pubSub := srv.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		logrus.Errorf("pubSub.Connect error: %v", err)
	}
	_, err := pubSub.Subscribe("user.create", func(publication broker.Publication) error {
		var user *pd.User
		if err := json.Unmarshal(publication.Message().Body, &user); err != nil {
			logrus.Errorf("json.Unmarshal error: %v", err)
			return err
		}
		logrus.Infoln("user is", user)
		go sendEmail(user)
		return nil
	})
	if err != nil {
		logrus.Errorf("pubSub.Subscribe error: %v", err)
	}

	// go-micro publish 注册监听事件
	micro.RegisterSubscriber("user.create", srv.Server(), new(EmailQueue))

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

func sendEmail(user *pd.User) error {
	log.Println("Sending email to:", user.Name)
	return nil
}

type EmailQueue struct{}

func (sub *EmailQueue) Process(ctx context.Context, user *pd.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}
