package main

import (
	"mygo/mongopool"
	"time"

	"fmt"

	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dbname = "yaoqu"
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

func SaveUser(p User) (string, error) {
	p.Id = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := mongopool.WithCollection(dbname, userCollName, query)
	return p.Id.Hex(), err
}

func RemoveById(id string) error { //dao层
	dbid := bson.ObjectIdHex(id)
	query := func(c *mgo.Collection) error {
		return c.RemoveId(dbid)
	}
	mongopool.WithCollection(dbname, userCollName, query)
	return nil
}

func Update(query bson.M, change bson.M) string {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := mongopool.WithCollection(dbname, userCollName, exop)
	if err != nil {
		return "true"
	}
	return "false"
}

func UpdateById(id string, change bson.M) string {
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.UpdateId(dbid, change)
	}
	err := mongopool.WithCollection(dbname, userCollName, exop)
	if err != nil {
		return "true"
	}
	return "false"
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func Find(query bson.M, fields bson.M, sort string, skip, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sort).Skip(skip).Limit(limit).All(&results)
	}
	mongopool.WithCollection(dbname, userCollName, exop)
	return
}
func FindById(id string, fields bson.M) (User, error) {
	dbid := bson.ObjectIdHex(id)
	user := User{}
	exop := func(c *mgo.Collection) error {
		return c.FindId(dbid).Select(fields).One(&user)
	}
	err := mongopool.WithCollection(dbname, userCollName, exop)
	return user, err
}

func GetUserById(id string) (User, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	user := User{}
	query := func(c *mgo.Collection) error {
		return c.FindId(dbid).One(&user)
	}
	err := mongopool.WithCollection(dbname, userCollName, query)
	return user, err
}

func GetAllUsers() ([]User, error) {
	var users []User
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&users)
	}
	err := mongopool.WithCollection(dbname, userCollName, query)
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetAllUsersCount() (int, error) {
	query := func(c *mgo.Collection) (int, error) {
		return c.Count()
	}
	return mongopool.WithCollection2(dbname, userCollName, query)
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

	//user := User{
	//	Nickname: "test11111",
	//}
	//u, err := SaveUser(user)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(u, err)
	//RemoveById("5b835d2328ddf5761c87fa04")

	//query := bson.M{"_id": bson.ObjectIdHex("5b835d8e28ddf57654b5ee16")}
	//res := Update(query, bson.M{"nickname": "测试"})
	//fmt.Println(res, "--------")
	//
	//res1 := UpdateById("5b835ef328ddf576fe94e65e", bson.M{"nickname": "测试222"})
	//fmt.Println(res1, "--------")

	//user, err := FindById("5b835ef328ddf576fe94e65e", bson.M{"nickname": 1, "_id": 1})
	//jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	//b, err := jsonIterator.Marshal(user) //encoding/json
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println(string(b), err)

	//user, err := Find(bson.M{"nickname": "test"}, bson.M{"nickname": 1, "_id": 1}, "nickname", 0, 10)
	//jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	//b, err := jsonIterator.Marshal(user) //encoding/json
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println(string(b), err)

	//users, err := GetAllUsers()
	//jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	//b, err := jsonIterator.Marshal(users) //encoding/json
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println(len(users), string(b))

	time.Sleep(1 * time.Second)

	for i := 0; i < 50; i++ {
		go func(i int) {
			c, err := GetAllUsersCount()
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(i, "---", c)
		}(i)
	}

	time.Sleep(1 * time.Second)

	cc := 1
	for i := 0; i < 200; i++ {
		go func(i int) {
			c, err := GetAllUsersCount()
			if err != nil {
				fmt.Println("error:", err)
			}
			cc++
			fmt.Println(i, "--again--", c)
			if cc == 200 {
				fmt.Println("done")
			}
		}(i)

	}
	time.Sleep(500 * time.Second)
}
