package mongopool

import (
	"fmt"

	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"gopkg.in/mgo.v2"
)

var (
	database = config.MongoDatabase
	username = config.MongoUsername
	password = config.MongoPassword
)

const limitConn = 50

var (
	mgoSessionChan chan *mgo.Session
	mgoSessions    []*mgo.Session
)

func init() {
	mgoSessionChan = make(chan *mgo.Session, 10000)
	mgoSessions = []*mgo.Session{}

	addr := config.MongodbClusterIP

	ssChanChan := make(chan chan *mgo.Session, limitConn*len(addr))
	go func() {
		for sessionCh := range ssChanChan {
			if session, ok := <-sessionCh; ok {
				mgoSessions = append(mgoSessions, session)
			}
		}
	}()

	for i := 0; i < limitConn; i++ {
		for _, host := range addr {
			ssChanChan <- createConnection(host)
		}
	}

	go func() {
		for {
			if len(mgoSessionChan) < 10000 {
				for _, s := range mgoSessions {
					if s != nil {
						mgoSessionChan <- s
					}
				}
			}
			time.Sleep(limitConn * time.Millisecond)
			//fmt.Println(len(mgoSessionChan), "--mgoSessionChan--")
		}

	}()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("init mongo connection to mgoSessionChan ...", len(mgoSessionChan))
	}()

}

func createConnection(host string) chan *mgo.Session {
	out := make(chan *mgo.Session)
	go func() {
		dialInfo := mgo.DialInfo{
			Addrs:    []string{host},
			Database: database,
			Username: username,
			Password: password,
			//PoolLimit: 50000,
			Timeout: time.Duration(60 * time.Second),
		}
		//fmt.Println("ing...",host)
		session, err := mgo.DialWithInfo(&dialInfo)
		//fmt.Println("done...",host)
		if err != nil || session == nil {
			fmt.Println(session, err)
			out <- nil
			return
		}
		session.SetMode(mgo.Monotonic, true)
		session.SetSafe(&mgo.Safe{
			WMode: "majority",
		})
		out <- session
	}()
	return out

}

func WithCollection(dbname, collName string, s func(*mgo.Collection) error) error {
	session := <-mgoSessionChan
	c := session.DB(dbname).C(collName)
	return s(c)
}

func WithCollection2(dbname, collName string, s func(*mgo.Collection) (int, error)) (int, error) {
	session := <-mgoSessionChan
	c := session.DB(dbname).C(collName)
	return s(c)
}
