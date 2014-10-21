package mongodb

import (
	"fmt"
	"labix.org/v2/mgo"
)

var dbSession *mgo.Session

func GetMongodbSession(hosts string) (*mgo.Session, error) {
	if dbSession == nil {
		fmt.Println("初始化数据库链接")
		var err error
		dbSession, err = mgo.Dial(hosts)
		if err != nil {
			return nil, err
		}
	}
	// 判断session是否出现错误
	if dbSession.Ping() != nil {
		fmt.Println("链接失效，重新链接")
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
