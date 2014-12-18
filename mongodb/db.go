package mongodb

import (
	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session

func GetMongodbSession(hosts string) (*mgo.Session, error) {
	if dbSession == nil {
		var err error
		dbSession, err = mgo.Dial(hosts)
		if err != nil {
			return nil, err
		}
	}
	if dbSession.Ping() != nil {
		var err error
		dbSession, err = mgo.Dial(hosts)
		if err != nil {
			return nil, err
		}
	}
	return dbSession.Clone(), nil
}
func GetCollections(name string, hosts string) *mgo.Collection {
	session, _ := GetMongodbSession(hosts)
	c := session.DB("wex").C(name)
	return c
}

func GetCollectionsWithDbName(db string, name string, hosts string) *mgo.Collection {
	session, _ := GetMongodbSession(hosts)
	c := session.DB(db).C(name)
	return c
}
