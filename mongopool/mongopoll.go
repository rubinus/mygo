package mongopool

import (
	"fmt"

	"time"

	"os"
	"yaosocket/config"

	"gopkg.in/mgo.v2"
)

var Addr = []string{
	"192.168.2.1:22001",
	"192.168.2.1:22002",
	"192.168.2.1:22003",
	//"127.0.0.1:27017",
}
var (
	database     = "yaoqu"
	username     = ""
	password     = ""
	mondbhostStr = ""
)

const limitConn = 10

var (
	mgoSessionChan = make(chan *mgo.Session)
	mgoSessions    []*mgo.Session
	mgoSessMap     = make(map[string]*mgo.Session)
)

func init() {
	go func() {
		for {
			for _, s := range mgoSessions {
				if s != nil {
					mgoSessionChan <- s
				}
			}
		}
	}()

	envHost := os.Getenv("MONGO_HOST")
	if envHost != "" {
		mondbhostStr = fmt.Sprintf("%s:%d", envHost, config.MongoDBPort)
	} else {
		mondbhostStr = fmt.Sprintf("%s:%d", config.DefaultMongoHost, config.MongoDBPort)
	}
	addr := []string{
		mondbhostStr,
	}
	if false {
		addr = Addr
	}
	for _, host := range addr {
		if _, ok := mgoSessMap[host]; !ok {
			fmt.Println("new session create ...", host)
			for i := 0; i < limitConn; i++ {
				createConnection(host)
			}
		}
	}
	//fmt.Println(len(mgoSessions))
}

func createConnection(host string) {
	dialInfo := mgo.DialInfo{
		Addrs:    []string{host},
		Database: database,
		Username: username,
		Password: password,
		Timeout:  time.Duration(15 * time.Second),
	}
	session, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		fmt.Println(session, err)
		return
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSafe(&mgo.Safe{
		WMode: "majority",
	})
	mgoSessMap[host] = session
	mgoSessions = append(mgoSessions, session)
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
