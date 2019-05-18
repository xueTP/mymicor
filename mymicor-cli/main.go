package main

import (
	"context"
	"fmt"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-server"
)

func main() {
	err := cmd.Init()
	client := pd.NewShippingServiceClient("go_micro_srv_consignment", microClient.DefaultClient)
	res, err := client.CreateConsignment(context.Background(), &pd.Consignment{Id: "123",})
	if err != nil {
		logrus.Errorf("ShippingServiceClient CreateConsignment error: %v", err)
	}
	fmt.Printf("%v, %v", res.Created, res.Consignment)
}
