package lib

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-micro/cmd"
	vesselPd "github.com/xueTP/gen-proto/mymicor-vessel"
	microClient "github.com/micro/go-micro/client"
)

type vesselLib struct{}

func NewVesselLib() vesselLib {
	return vesselLib{}
}

type VesselLibInterface interface{
	FindAvailable(capacity, maxWeight int32) *vesselPd.Vessel
}

func getVesselClient() vesselPd.VesselServiceClient {
	cmd.Init()
	client := vesselPd.NewVesselServiceClient("go_micro_srv_vessel", microClient.DefaultClient)
	return client
}

func (this vesselLib) FindAvailable(capacity, maxWeight int32) *vesselPd.Vessel {
	client := getVesselClient()
	specification := &vesselPd.Specification{Capacity: capacity, MaxWeight: maxWeight}
	res, err := client.FindAvailable(context.Background(), specification)
	if err != nil {
		logrus.Errorf("client FindAvailable error: %v, param: %+v", err, specification)
		return &vesselPd.Vessel{}
	}
	return res.Vessel
}