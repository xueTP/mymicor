package data

import "gopkg.in/mgo.v2"

func CreateSession(host string) (*mgo.Session, error) {
	sess, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	sess.SetMode(mgo.Monotonic, true)
	return sess, nil
}
