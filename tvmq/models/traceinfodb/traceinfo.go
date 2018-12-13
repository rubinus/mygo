package traceinfodb

import (
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/mongopool"
	"github.com/json-iterator/go"
	"github.com/kataras/iris/core/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var database = config.MongoDatabase
var TraceinfoCollName = "traceinfo"

type Traceinfo struct {
	Id                   bson.ObjectId `json:"id,omitempty" bson:"_id"`
	ErrMsg               string        `json:"errMsg,omitempty" bson:"errMsg,omitempty"`
	Ttype                string        `json:"tType,omitempty" bson:"tType,omitempty"`
	TraceId              string        `json:"traceId,omitempty" bson:"traceId,omitempty"`
	TenantId             string        `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	UserId               string        `json:"userId,omitempty" bson:"userId,omitempty"`
	RecvByClient         int64         `json:"recvByClient,omitempty" bson:"recvByClient,omitempty"`
	FilterWord           int64         `json:"filterWord,omitempty" bson:"filterWord,omitempty"`
	GetUserInfo          int64         `json:"getUserInfo,omitempty"  bson:"getUserInfo,omitempty"`
	SaveUserToHashAndSet int64         `json:"saveUserToHashAndSet,omitempty"  bson:"saveUserToHashAndSet,omitempty"`
	CheckIsActivity      int64         `json:"checkIsActivity,omitempty"  bson:"checkIsActivity,omitempty"`
	MsgLimit             int64         `json:"msgLimit,omitempty"  bson:"msgLimit,omitempty"`
	RecvByQueue          int64         `json:"recvByQueue,omitempty" bson:"recvByQueue,omitempty"`
	RecvByNode           int64         `json:"recvByNode,omitempty" bson:"recvByNode,omitempty"`
	RecvByCallGift       int64         `json:"recvByCallGift,omitempty" bson:"recvByCallGift,omitempty"`
	RecvBySaveMongo      int64         `json:"recvBySaveMongo,omitempty" bson:"recvBySaveMongo,omitempty"`
	CallRPCClient        int64         `json:"callRPCClient,omitempty" bson:"callRPCClient,omitempty"`
	RecvByRPCServe       int64         `json:"recvByRPCServe,omitempty" bson:"recvByRPCServe,omitempty"`
	SendToClient         int64         `json:"sendToClient,omitempty" bson:"sendToClient,omitempty"`
	CreateTime           int64         `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime           int64         `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

func (u *Traceinfo) StructToMap() map[string]interface{} {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	var m map[string]interface{}
	b, _ := jsonIterator.Marshal(u)
	jsonIterator.Unmarshal(b, &m)
	return m
}

func (p *Traceinfo) SaveTraceinfo() (string, error) {
	p.Id = bson.NewObjectId()
	if p.CreateTime == 0 {
		p.CreateTime = time.Now().UnixNano() / 1e6
	}
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := mongopool.WithCollection(database, TraceinfoCollName, query)
	return p.Id.Hex(), err
}

func (p *Traceinfo) RemoveById(id string) error { //dao层
	dbid := bson.ObjectIdHex(id)
	query := func(c *mgo.Collection) error {
		return c.RemoveId(dbid)
	}
	mongopool.WithCollection(database, TraceinfoCollName, query)
	return nil
}

func (p *Traceinfo) Update(query bson.M, change bson.M) error {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := mongopool.WithCollection(database, TraceinfoCollName, exop)
	if err != nil {
		return err
	}
	return nil
}

func (p *Traceinfo) UpdateById(id string, change bson.M) error {
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.UpdateId(dbid, change)
	}
	err := mongopool.WithCollection(database, TraceinfoCollName, exop)
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
func (p *Traceinfo) Find(query bson.M, fields bson.M, sort string, skip, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sort).Skip(skip).Limit(limit).All(&results)
	}
	mongopool.WithCollection(database, TraceinfoCollName, exop)
	return
}
func (p *Traceinfo) FindById(id string, fields bson.M) (Traceinfo, error) {
	traceinfo := Traceinfo{}
	if !bson.IsObjectIdHex(id) {
		return traceinfo, errors.New("not object id hex")
	}
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.FindId(dbid).Select(fields).One(&traceinfo)
	}
	err := mongopool.WithCollection(database, TraceinfoCollName, exop)
	return traceinfo, err
}

func (p *Traceinfo) GetTraceinfoById(id string) (Traceinfo, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	traceinfo := Traceinfo{}
	query := func(c *mgo.Collection) error {
		return c.FindId(dbid).One(&traceinfo)
	}
	err := mongopool.WithCollection(database, TraceinfoCollName, query)
	return traceinfo, err
}

func (p *Traceinfo) GetTraceinfoByTraceId() (*Traceinfo, error) { //dao层
	traceinfo := &Traceinfo{}
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"traceId": p.TraceId}).One(traceinfo)
	}
	err := mongopool.WithCollection(database, TraceinfoCollName, query)
	return traceinfo, err
}
