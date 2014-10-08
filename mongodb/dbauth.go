package mongodb

import (
	"fmt"
	"labix.org/v2/mgo"
	"strings"
)

var dbAuthSession map[string]*mgo.Session

func GetAuthMongodbSession(hosts string, user string, password string, db string) (*mgo.Session, error) {
	if dbAuthSession == nil {
		dbAuthSession = make(map[string]*mgo.Session)
	}
	s := dbAuthSession[db]
	if s == nil {
		fmt.Println("初始化数据库链接")
		var err error
		dial := new(mgo.DialInfo)
		fmt.Println(strings.Split(hosts, ","))
		dial.Addrs = strings.Split(hosts, ",")
		dial.Username = user
		dial.Password = password
		dial.Database = db
		s, err = mgo.DialWithInfo(dial)
		if err != nil {
			return nil, err
		}
		dbAuthSession[db] = s

	}
	return s.Clone(), nil
}
func GetAuthCollectionsWithName(name string, db string, session *mgo.Session) *mgo.Collection {
	c := session.DB(db).C(name)
	return c
}
