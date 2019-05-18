package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro"
	pd "github.com/xueTP/gen-proto/mymicor-consignment"
	"mymicor/mymicor-server/data"
	"mymicor/mymicor-server/server"
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
		// This name must match the package name given in your protobuf definition
		micro.Name("go_micro_srv_consignment"),
		micro.Version("latest"),
	)
	// Init will parse the command line flags.
	srv.Init()
	// Get Server Handle
	consignmentServer := server.NewConsignmentServer(session)
	// Register handler
	pd.RegisterShippingServiceHandler(srv.Server(), consignmentServer)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}