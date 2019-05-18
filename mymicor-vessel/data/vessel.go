package data

import (
	"gopkg.in/mgo.v2"
	pd "github.com/xueTP/gen-proto/mymicor-vessel"
)

const (
	DBNAME = "shippy"
	COLLECTION = "vessel"
)

type VesselDataer interface {
	GetAll() ([]*pd.Vessel, error)
	Close()
}

type vesselData struct {
	session *mgo.Session
}

func NewVesselData(session *mgo.Session) vesselData {
	return vesselData{
		session: session,
	}
}

func (this vesselData) collection() *mgo.Collection {
	return this.session.DB(DBNAME).C(COLLECTION)
}

func (this vesselData) GetAll() ([]*pd.Vessel, error) {
	res := []*pd.Vessel{}
	err := this.collection().Find(nil).All(&res)
	return res, err
}

func (this vesselData) Close() {
	this.session.Close()
}
