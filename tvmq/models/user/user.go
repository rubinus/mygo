package user

import (
	"fmt"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/mongopool"
	"github.com/json-iterator/go"
	"github.com/kataras/iris/core/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var database = config.MongoDatabase
var userCollName = "friends"

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Unionid    string        `json:"unionid,omitempty" bson:"unionid,omitempty"`
	MinAppid   string        `json:"minappid" bson:"minappid"`
	MinOpenid  string        `json:"minopenid" bson:"minopenid"`
	Nickname   string        `json:"nickname,omitempty" bson:"nickname,omitempty"`
	Headimgurl string        `json:"headimgurl,omitempty" bson:"headimgurl,omitempty"`
	City       string        `json:"city,omitempty" bson:"city,omitempty"`
	Sex        string        `json:"sex,omitempty" bson:"sex,omitempty"`
	Province   string        `json:"province,omitempty" bson:"province,omitempty"`
	Country    string        `json:"country,omitempty" bson:"country,omitempty"`
	IsAi       string        `json:"isAi,omitempty" bson:"isAi,omitempty"`
	CreateTime int64         `json:"createTime" bson:"createTime"`
	UpdateTime int64         `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

func SaveUser(p *User) (string, error) {
	p.Id = bson.NewObjectId()
	if p.CreateTime == 0 {
		p.CreateTime = time.Now().UnixNano() / 1e6
	}
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := mongopool.WithCollection(database, userCollName, query)
	return p.Id.Hex(), err
}

func RemoveById(id string) error { //dao层
	dbid := bson.ObjectIdHex(id)
	query := func(c *mgo.Collection) error {
		return c.RemoveId(dbid)
	}
	mongopool.WithCollection(database, userCollName, query)
	return nil
}

func Update(query bson.M, change bson.M) error {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := mongopool.WithCollection(database, userCollName, exop)
	if err != nil {
		return err
	}
	return nil
}

func UpdateById(id string, change bson.M) error {
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.UpdateId(dbid, change)
	}
	err := mongopool.WithCollection(database, userCollName, exop)
	if err != nil {
		return err
	}
	return nil
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
	mongopool.WithCollection(database, userCollName, exop)
	return
}
func FindById(id string, fields bson.M) (User, error) {
	user := User{}
	if !bson.IsObjectIdHex(id) {
		return user, errors.New("not object id hex")
	}
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.FindId(dbid).Select(fields).One(&user)
	}
	err := mongopool.WithCollection(database, userCollName, exop)
	return user, err
}

func GetUserById(id string) (User, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	user := User{}
	query := func(c *mgo.Collection) error {
		return c.FindId(dbid).One(&user)
	}
	err := mongopool.WithCollection(database, userCollName, query)
	return user, err
}

func FindByAppidOpenid(appid, openid string) (*User, error) {
	var user *User
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"minappid": appid, "minopenid": openid}).One(&user)
	}
	err := mongopool.WithCollection(database, userCollName, query)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAllUsers() ([]User, error) {
	fmt.Println(database, userCollName)
	var users []User
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&users)
	}
	err := mongopool.WithCollection(database, userCollName, query)
	if err != nil {
		return users, err
	}
	return users, nil
}

func getUserinfo(i int) {
	user, err1 := GetUserById("5ae01b2d38e89043f505ddb3")
	jsonIterator1 := jsoniter.ConfigCompatibleWithStandardLibrary
	b1, err1 := jsonIterator1.Marshal(user) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(i, "======", string(b1))

}
