package mongodb

import (
	"gopkg.in/mgo.v2"
	"strings"
)

var dbAuthSession map[string]*mgo.Session

func GetAuthMongodbSession(hosts string, user string, password string, db string) (*mgo.Session, error) {
	if dbAuthSession == nil {
		dbAuthSession = make(map[string]*mgo.Session)
		var err error
		dial := new(mgo.DialInfo)
		dial.Addrs = strings.Split(hosts, ",")
		dial.Username = user
		dial.Password = password
		dial.Database = db
		ss, err := mgo.DialWithInfo(dial)
		if err != nil {
			panic(err)
			return nil, err
		}
		dbAuthSession[db] = ss
		return ss.Clone(), nil
	}
	s := dbAuthSession[db]
	if s == nil {
		goto GETSESSION
	} else {
		if s.Ping() != nil {
			return s.Clone(), nil
		}
		goto GETSESSION
	}
GETSESSION:
	var err error
	dial := new(mgo.DialInfo)
	dial.Addrs = strings.Split(hosts, ",")
	dial.Username = user
	dial.Password = password
	dial.Database = db
	ss, err := mgo.DialWithInfo(dial)
	if err != nil {
		panic(err)
		return nil, err
	}
	dbAuthSession[db] = ss
	return ss.Clone(), nil

}

/*
	func GetAuthMongodbSession(hosts string, user string, password string, db string) (*mgo.Session, error) {
		if dbAuthSession == nil {
			dbAuthSession = make(map[string]*mgo.Session)
		}
		s := dbAuthSession[db]
		if s == nil {
			var err error
			dial := new(mgo.DialInfo)
			dial.Addrs = strings.Split(hosts, ",")
			dial.Username = user
			dial.Password = password
			dial.Database = db
			s, err = mgo.DialWithInfo(dial)
			if err != nil {
				panic(err)
				return nil, err
			}
			dbAuthSession[db] = s

		}
		return s.Clone(), nil
	}
*/
func GetAuthCollectionsWithName(name string, db string, session *mgo.Session) *mgo.Collection {
	c := session.DB(db).C(name)
	return c
}
