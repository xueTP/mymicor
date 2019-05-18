package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/micro/go-micro"
	"mymicor/mymicor-user/data"
	"mymicor/mymicor-user/server"
	pd "github.com/xueTP/gen-proto/mymicor-user"
)

func main() {
	// gorm create conn
	conn, err := data.CreateConnection()
	if err != nil {
		logrus.Errorf("create gorm connection error: %v", err)
	}
	defer conn.Close()
	// server init
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	srv.Init()

	userServer := server.NewUserServer(conn)
	// Register handler
	pd.RegisterUserServiceHandler(srv.Server(), userServer)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
