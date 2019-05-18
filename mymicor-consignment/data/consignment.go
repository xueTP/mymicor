package data

import (
	"gopkg.in/mgo.v2"
	pd "github.com/xueTP/gen-proto/mymicor-consignment"
)

const (
	DBNAME = "shippy"
	COLLECTION = "consignment"
)

type ConsignmentDataer interface {
	Create(consignment *pd.Consignment) error
	GetAll() ([]*pd.Consignment, error)
	Close()
}

type consignmentData struct {
	session *mgo.Session
}

func NewConsignmentData(session *mgo.Session) consignmentData {
	return consignmentData{
		session: session,
	}
}

func (this consignmentData) collection() *mgo.Collection {
	return this.session.DB(DBNAME).C(COLLECTION)
}

func (this consignmentData) Create(consignment *pd.Consignment) error {
	return this.collection().Insert(consignment)
}

func (this consignmentData) GetAll() ([]*pd.Consignment, error) {
	res := []*pd.Consignment{}
	err := this.collection().Find(nil).All(&res)
	return res, err
}

func (this consignmentData) Close() {
	this.session.Close()
}
