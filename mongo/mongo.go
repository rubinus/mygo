package mongo

import (
	"fmt"
	"log"

	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestMongo() {
	session, err := mgo.Dial("106.15.228.49:27027")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("intggnet").C("users")
	type Person struct {
		Nickname string `json:"nickname"`
		Country  string `json:"country"`
	}
	result := Person{}
	err = c.Find(bson.M{"nickname": "林"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n--result-------------------", result)
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err1 := jsonIterator.Marshal(result) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", b)
	//os.Stdout.Write(b)
	fmt.Println("\n--result-------------------\n")

	fmt.Printf("mongodb======Name: %s,=== country== %s\n", result.Country, result.Nickname)
}
