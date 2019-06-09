package server

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	pd "github.com/xueTP/gen-proto/mymicor-vessel"
	"gopkg.in/mgo.v2"
	"mymicor/mymicor-vessel/data"
)

type vesselServer struct {
	session *mgo.Session
	vesselData data.VesselDataer
}

func NewVesselServer(session *mgo.Session) vesselServer {
	return vesselServer{
		session: session,
		vesselData: data.NewVesselData(session.Clone()),
	}
}

func (this vesselServer) FindAvailable(ctx context.Context,spec *pd.Specification, res *pd.Response) (error) {
	defer this.session.Clone()
	logrus.Infof("start into CreateConsignment, %v", spec)
	vessels, err := this.vesselData.GetAll()
	if err != nil {
		logrus.Errorf("vesselServer.FindAvailable error: %v", err)
		return err
	}
	for _, v := range vessels {
		if spec.Capacity <= v.Capacity && spec.MaxWeight <= v.MaxWeight {
			res.Vessel = v
			return nil
		}
	}
	return errors.New("not find vessel")
}
