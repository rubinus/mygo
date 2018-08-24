package main

import (
	"mygo/mongopool"
	"time"

	"fmt"

	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2/bson"
)

var userCollName = "users"

type User struct {
	Id         bson.ObjectId `bson:"_id"`
	Nickname   string        `json:"nickname,omitempty"`
	Headimgurl string        `json:"headimgurl,omitempty"`
	City       string        `json:"city,omitempty"`
	Province   string        `json:"province,omitempty"`
	Country    string        `json:"country,omitempty"`
	CreateTime *time.Time    `json:"createTime"`
	UpdateTime *time.Time    `json:"updateTime,omitempty"`
}

func GetUserById(id string) (User, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	user := User{}
	query := func(c *mongopool.Coll) error {
		return c.FindId(dbid).One(&user)
	}
	mongopool.CloserColl(userCollName, query)
	return user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	query := func(c *mongopool.Coll) error {
		return c.Find(nil).All(&users)
	}
	err := mongopool.CloserColl(userCollName, query)
	if err != nil {
		return users, nil
	}
	return users, nil
}

func getUserinfo(i int) {

	//c, err := mongopool.GetColl("users")
	//if err != nil {
	//	panic(err)
	//}
	//defer c.Close()
	//
	//result := User{}
	//id := "5ae01b2d38e89043f505ddb3"
	//err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//
	//jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	//b, err := jsonIterator.Marshal(result) //encoding/json
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println(i, "-----", string(b))

	user, err1 := GetUserById("5ae01b2d38e89043f505ddb3")
	jsonIterator1 := jsoniter.ConfigCompatibleWithStandardLibrary
	b1, err1 := jsonIterator1.Marshal(user) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(i, "======", string(b1))

}

func main() {

	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		getUserinfo(i)
	//	}(i)
	//}
	//
	//time.Sleep(1 * time.Second)
	//
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		getUserinfo(i)
	//	}(i)
	//}
	//
	//time.Sleep(5 * time.Second)

	users, err := GetAllUsers()
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := jsonIterator.Marshal(users) //encoding/json
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(len(users), string(b))

}
