package gift

import (
	"fmt"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/mongopool"
	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var database = config.MongoDatabase
var giftCollName = "gifts"

type Gift struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Appid      string        `json:"appid" bson:"appid"`
	Userid     string        `json:"userid" bson:"userid"`
	TraceId    string        `json:"traceId,omitempty" bson:"traceId,omitempty"`
	Nickname   string        `json:"nickname,omitempty" bson:"nickname"`
	Headimgurl string        `json:"headimgurl,omitempty" bson:"headimgurl"`
	Giftid     string        `json:"giftid,omitempty" bson:"giftid"`
	Giftname   string        `json:"giftname,omitempty" bson:"giftname"`
	Icon       string        `json:"icon,omitempty" bson:"icon"`
	Pictures   string        `json:"pictures,omitempty" bson:"pictures"`
	Count      int           `json:"count,omitempty" bson:"count"`
	Points     int           `json:"points,omitempty" bson:"points"`
	CreateTime int64         `json:"createTime,omitempty" bson:"createTime"`
}

func (p Gift) SaveGift() (string, error) {
	p.Id = bson.NewObjectId()
	if p.CreateTime == 0 {
		p.CreateTime = time.Now().UnixNano() / 1e6
	}
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := mongopool.WithCollection(database, giftCollName, query)
	return p.Id.Hex(), err
}

func (p Gift) RemoveById(id string) error { //dao层
	dbid := bson.ObjectIdHex(id)
	query := func(c *mgo.Collection) error {
		return c.RemoveId(dbid)
	}
	mongopool.WithCollection(database, giftCollName, query)
	return nil
}

func (p Gift) Update(query bson.M, change bson.M) string {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := mongopool.WithCollection(database, giftCollName, exop)
	if err != nil {
		return "true"
	}
	return "false"
}

func (p Gift) UpdateById(id string, change bson.M) string {
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.UpdateId(dbid, change)
	}
	err := mongopool.WithCollection(database, giftCollName, exop)
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
func (p Gift) Find(query bson.M, fields bson.M, sort string, skip, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sort).Skip(skip).Limit(limit).All(&results)
	}
	mongopool.WithCollection(database, giftCollName, exop)
	return
}
func (p Gift) FindById(id string, fields bson.M) (Gift, error) {
	dbid := bson.ObjectIdHex(id)
	gift := Gift{}
	exop := func(c *mgo.Collection) error {
		return c.FindId(dbid).Select(fields).One(&gift)
	}
	err := mongopool.WithCollection(database, giftCollName, exop)
	return gift, err
}

func (p Gift) GetgiftById(id string) (Gift, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	gift := Gift{}
	query := func(c *mgo.Collection) error {
		return c.FindId(dbid).One(&gift)
	}
	err := mongopool.WithCollection(database, giftCollName, query)
	return gift, err
}

func (p Gift) GetAllgifts() ([]Gift, error) {
	var gifts []Gift
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&gifts)
	}
	err := mongopool.WithCollection(database, giftCollName, query)
	if err != nil {
		return gifts, err
	}
	return gifts, nil
}

func getGiftinfo(i int) {
	var g Gift
	gift, err1 := g.GetgiftById("5ae01b2d38e89043f505ddb3")
	jsonIterator1 := jsoniter.ConfigCompatibleWithStandardLibrary
	b1, err1 := jsonIterator1.Marshal(gift) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(i, "======", string(b1))

}
