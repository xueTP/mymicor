package server

import (
	"context"
	"github.com/Sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-consignment"
	"gopkg.in/mgo.v2"
	"mymicor/mymicor-server/data"
	"mymicor/mymicor-server/lib"
)

type ConsignmentServer struct{
	session *mgo.Session
	consignmentData data.ConsignmentDataer
	VesselLib lib.VesselLibInterface
}

var consignmentd ConsignmentServer

func NewConsignmentServer(s *mgo.Session) ConsignmentServer {
	consignmentd = ConsignmentServer{
		session: s,
		VesselLib: lib.NewVesselLib(),
	}
	consignmentd.consignmentData = data.NewConsignmentData(consignmentd.session.Clone())
	return consignmentd
}

func (cs ConsignmentServer) CreateConsignment (ctx context.Context, consignment *pd.Consignment, res *pd.Response) (error) {
	defer cs.consignmentData.Close()
	logrus.Infoln("start into CreateConsignment")
	vessel := cs.VesselLib.FindAvailable(int32(len(consignment.Containers)), consignment.Weight)
	consignment.VesselId = vessel.Id
	err := cs.consignmentData.Create(consignment)
	if err != nil {
		logrus.Errorf("ConsignmentServer.CreateConsignment error: %v, param: %+v", err, consignment)
		return err
	}
	logrus.Infoln("end out CreateConsignment")
	return nil
}

func (cs ConsignmentServer) GetConsignments(ctx context.Context, req *pd.GetRequest, res *pd.Response) error {
	defer cs.consignmentData.Close()
	var err error
	res.Consignments, err = cs.consignmentData.GetAll()
	if err != nil {
		logrus.Errorf("ConsignmentServer.GetConsignments error: %v", err)
		return err
	}
	return nil
}
