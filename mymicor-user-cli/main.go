package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-user"
	"golang.org/x/net/context"
	"log"
	"os"
)

func main(){
	err := cmd.Init()
	if err != nil {
		panic(err)
	}
	// user client
	client := pd.NewUserServiceClient("go.micro.srv.user", microClient.DefaultClient)
	// cli server flag
	server := micro.NewService(
		micro.Flags(
			cli.StringFlag{Name: "name", Usage: "you full name"},
			cli.StringFlag{Name: "email", Usage: "you email"},
			cli.StringFlag{Name: "password", Usage: "you password"},
			cli.StringFlag{Name: "company", Usage: "you company"},
		),
	)
	
	server.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")
			if name == "" {
				name = "xue"
				email = "xue@qq.com"
				password = "123456"
				company = "xxx.ltd"
			}
			user := pd.User{Name: name, Email: email, Password: password, Company: company}
			fmt.Println("#############", user)

			r, err := client.Create(context.TODO(), &user)
			if err != nil {
				logrus.Errorf("user server client create err: %v, user: %+v", err, user)
			}
			logrus.Printf("create user success, user: %+v", r.User)
			token, err := client.Auth(context.TODO(), &pd.User{Email: email, Password: password})
			if err != nil || token == nil {
				log.Fatalf("user server Auth error: %v, token: %v", err, token)
			}
			logrus.Println("token is", token.Token)
			ctx := metadata.NewContext(context.Background(), map[string]string{"token": token.Token})
			getAll, err := client.GetAll(ctx, &pd.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),	
	)

	// 运行
	if err := server.Run(); err != nil {
		log.Println(err)
	}
}
