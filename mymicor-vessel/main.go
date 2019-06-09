package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-vessel"
	"mymicor/mymicor-vessel/data"
	"mymicor/mymicor-vessel/server"
	"os"
)

const (
	DEFAULT_DB_HOST = ":27017"
)

func main() {
	// get mongodb session
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_DB_HOST
	}
	session, err := data.CreateSession(dbHost)
	if err != nil {
		logrus.Errorf("get mongodb session error: %v, host: %s", err, dbHost)
		return
	}
	defer session.Close()

	srv := micro.NewService(
		micro.Name("go_micro_srv_vessel"),
		micro.Version("latest"),
	)
	srv.Init()

	vesselServer := server.NewVesselServer(session)
	// Register handler
	pd.RegisterVesselServiceHandler(srv.Server(), vesselServer)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
