package mongopool

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var Sessions = map[string]*mgo.Session{}

func NewSession(dbhost string) (*mgo.Session, error) {
	//if _, ok := Sessions[dbhost]; !ok {
	//	info := mgo.DialInfo{
	//		Addrs:     []string{dbhost},
	//		Source:    "admin",    // 设置权限的数据库 authdb: admin
	//		Timeout:   5 * time.Second,
	//		PoolLimit: 4096,
	//	}
	//
	//	fmt.Println(dbhost)
	//
	//	session, err = mgo.DialWithInfo(&info)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	session.SetMode(mgo.Monotonic, true)
	//	session.SetSafe(&mgo.Safe{
	//		WMode: "majority",
	//	})
	//	Sessions[dbhost] = session
	//}

	if _, ok := Sessions[dbhost]; !ok {
		fmt.Println("new session create ...")
		session, err := mgo.Dial(dbhost)
		if err != nil {
			return nil, err
		}
		session.SetMode(mgo.Monotonic, true)
		session.SetSafe(&mgo.Safe{
			WMode: "majority",
		})
		Sessions[dbhost] = session
	} else {
		fmt.Println("old session ...")
	}
	return Sessions[dbhost].Copy(), nil
}

type MongoDatabase struct {
	*mgo.Database
}

func NewDatabase(dbhost, dbName string) (*MongoDatabase, error) {
	session, err := NewSession(dbhost)
	if err != nil {
		return nil, err
	}
	return &MongoDatabase{session.DB(dbName)}, nil
}

type Coll struct {
	*mgo.Collection
}

func NewColl(dbhost, dbName, collName string) (*Coll, error) {
	db, err := NewDatabase(dbhost, dbName)
	if err != nil {
		return nil, err
	}
	return &Coll{db.C(collName)}, nil
}

func (m *Coll) Close() {
	m.Database.Session.Close()
}

func CloserColl(collName string, s func(*Coll) error) error {
	c, err := GetColl(collName)
	if err != nil {
		return err
	}
	defer c.Close()
	return s(c)
}

func GetColl(collName string) (*Coll, error) {
	dbhost := "106.15.228.49:27027"

	coll, err := NewColl(dbhost, "*******", collName)
	if err != nil {
		return nil, err
	}

	return coll, nil
}
