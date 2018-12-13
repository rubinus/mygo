package comment

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
var CommentCollName = "comments"

type Comment struct {
	Id         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Appid      string        `json:"appid" bson:"appid"`
	Userid     string        `json:"userid" bson:"userid"`
	TraceId    string        `json:"traceId,omitempty" bson:"traceId,omitempty"`
	TenantId   string        `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	Nickname   string        `json:"nickname,omitempty" bson:"nickname,omitempty"`
	Headimgurl string        `json:"headimgurl,omitempty" bson:"headimgurl,omitempty"`
	Content    string        `json:"content" bson:"content"`
	CreateTime int64         `json:"createTime,omitempty" bson:"createTime,omitempty"`
}

func (p Comment) SaveComment() (string, error) {
	p.Id = bson.NewObjectId()
	if p.CreateTime == 0 {
		p.CreateTime = time.Now().UnixNano() / 1e6
	}
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := mongopool.WithCollection(database, CommentCollName, query)
	return p.Id.Hex(), err
}

func (p Comment) RemoveById(id string) error { //dao层
	dbid := bson.ObjectIdHex(id)
	query := func(c *mgo.Collection) error {
		return c.RemoveId(dbid)
	}
	mongopool.WithCollection(database, CommentCollName, query)
	return nil
}

func (p Comment) Update(query bson.M, change bson.M) string {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := mongopool.WithCollection(database, CommentCollName, exop)
	if err != nil {
		return "true"
	}
	return "false"
}

func (p Comment) UpdateById(id string, change bson.M) string {
	dbid := bson.ObjectIdHex(id)
	exop := func(c *mgo.Collection) error {
		return c.UpdateId(dbid, change)
	}
	err := mongopool.WithCollection(database, CommentCollName, exop)
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
func (p Comment) Find(query bson.M, fields bson.M, sort string, skip, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sort).Skip(skip).Limit(limit).All(&results)
	}
	mongopool.WithCollection(database, CommentCollName, exop)
	return
}
func (p Comment) FindById(id string, fields bson.M) (Comment, error) {
	dbid := bson.ObjectIdHex(id)
	Comment := Comment{}
	exop := func(c *mgo.Collection) error {
		return c.FindId(dbid).Select(fields).One(&Comment)
	}
	err := mongopool.WithCollection(database, CommentCollName, exop)
	return Comment, err
}

func (p Comment) GetCommentById(id string) (Comment, error) { //dao层
	dbid := bson.ObjectIdHex(id)
	Comment := Comment{}
	query := func(c *mgo.Collection) error {
		return c.FindId(dbid).One(&Comment)
	}
	err := mongopool.WithCollection(database, CommentCollName, query)
	return Comment, err
}

func (p Comment) GetAllComments() ([]Comment, error) {
	var Comments []Comment
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&Comments)
	}
	err := mongopool.WithCollection(database, CommentCollName, query)
	if err != nil {
		return Comments, err
	}
	return Comments, nil
}

func getComment(i int) {
	var c Comment
	Comment, err1 := c.GetCommentById("5ae01b2d38e89043f505ddb3")
	jsonIterator1 := jsoniter.ConfigCompatibleWithStandardLibrary
	b1, err1 := jsonIterator1.Marshal(Comment) //encoding/json
	if err1 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(i, "======", string(b1))

}
