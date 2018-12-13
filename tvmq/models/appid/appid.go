package appid

import (
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/mongopool"
	"github.com/kataras/iris/core/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var database = config.MongoDatabase
var userCollName = "appids"

type Appid struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Appid      string        `json:"appid,omitempty" bson:"appid,omitempty"`
	Secret     string        `json:"secret" bson:"secret"`
	CreateTime int64         `json:"createTime" bson:"createTime"`
}

func SaveAppid(p Appid) (string, error) {
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
func FindById(id string, fields bson.M) (Appid, error) {
	appid := Appid{}
	if !bson.IsObjectIdHex(id) {
		return appid, errors.New("not object id hex")
	}
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.FindId(dbid).Select(fields).One(&appid)
	}
	err := mongopool.WithCollection(database, userCollName, exop)
	return appid, err
}

func GetUserById(id string) (Appid, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	appid := Appid{}
	query := func(c *mgo.Collection) error {
		return c.FindId(dbid).One(&appid)
	}
	err := mongopool.WithCollection(database, userCollName, query)
	return appid, err
}

func GetAllAppids() ([]Appid, error) {
	var appids []Appid
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&appids)
	}
	err := mongopool.WithCollection(database, userCollName, query)
	if err != nil {
		return appids, err
	}
	return appids, nil
}
