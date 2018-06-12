package mongo

import (
	"fmt"
	"log"

	"os"

	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestMongo() string {
	monHost := os.Getenv("MONGO_HOST")
	monStr := ""
	if monHost != "" {
		monStr = fmt.Sprintf("%s:%d", monHost, 27017)
	} else {
		monStr = fmt.Sprintf("%s:%d", "106.15.228.49", 27027)
	}
	session, err := mgo.Dial(monStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("local").C("startup_log")
	type Person struct {
		Hostname         string `json:"hostname"`
		Openssl          string `json:"openssl"`
		CmdLine          string `json:"cmdLine"`
		StorageEngines   string `json:"storageEngines"`
		BuildEnvironment string `json:"buildEnvironment"`
	}
	result := Person{}
	err = c.Find(bson.M{}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n--result-------------------", result)
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err1 := jsonIterator.Marshal(result) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err)
	}

	return string(b)
}
