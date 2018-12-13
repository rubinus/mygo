package routes

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"code.tvmining.com/tvplay/tvmq/allmap"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/lib"
	"code.tvmining.com/tvplay/tvmq/minauth"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/user"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/bitly/go-simplejson"
	"github.com/json-iterator/go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
)

type WxauthRequestBody struct {
	appid  string
	code   string
	secret string
}
type WxUserinfo struct {
	OpenId    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
	Watermark WxUserinfoMark
}
type WxUserinfoMark struct {
	Timestamp int64  `json:"timestamp"`
	Appid     string `json:"appid"`
}

type MapOpenidUserid struct {
	Userid  string `json:"userid"`
	Unionid string `json:"unionid,omitempty"`
}

type RedisUser struct {
	Id         string `json:"id,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	MinAppid   string `json:"minappid,omitempty"`
	MinOpenid  string `json:"minopenid,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Headimgurl string `json:"headimgurl,omitempty"`
	City       string `json:"city,omitempty"`
	Sex        string `json:"sex,omitempty"`
	Province   string `json:"province,omitempty"`
	Country    string `json:"country,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
	SessionKey string `json:"sessionKey,omitempty"`
	Token      string `json:"token,omitempty"`
}

func (u *RedisUser) StructToMap() map[string]interface{} {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	var m map[string]interface{}
	b, _ := jsonIterator.Marshal(u)
	jsonIterator.Unmarshal(b, &m)
	return m
}

func (u *MapOpenidUserid) StructToMap() map[string]interface{} {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	var m map[string]interface{}
	b, _ := jsonIterator.Marshal(u)
	jsonIterator.Unmarshal(b, &m)
	return m
}

func ChcekTokenRoute(ctx iris.Context) {
	userid := ctx.URLParam("userid")
	token := ctx.URLParam("token")
	//check 参数
	if userid == "" || token == "" {
		ctx.JSONP(iris.Map{"status": 401, "msg": "缺少必传参数userid或token"})
		return
	}
	//check redis token by user
	var redisUser *RedisUser
	key := fmt.Sprintf("%s:%s", config.PreixForUser, userid)
	if redisUserOld, err := CheckToken(key, token); err != nil {
		ctx.JSONP(iris.Map{"status": 401, "msg": err.Error()})
		return
	} else {
		redisUser = redisUserOld
	}
	if redisUser.Token == token {
		ctx.JSONP(iris.Map{"status": 200, "msg": "token is ok"})

	} else {
		ctx.JSONP(iris.Map{"status": 500, "msg": "token is invalid"})
	}
}

func GetUserinfo(ctx iris.Context) {
	userid := ctx.URLParam("userid")

	redisUser, err := GetUserinRedisOrMongo(userid)

	var result *simplejson.Json
	redisUser.Token = ""
	redisUser.SessionKey = ""

	r, _ := models.Marshal(redisUser)
	result, _ = simplejson.NewJson([]byte(r))
	if err != nil {
		ctx.JSONP(iris.Map{"status": 201, "msg": "no user"})
	} else {
		ctx.JSONP(iris.Map{"status": 200, "data": result})
	}
}

func GetUserinRedisOrMongo(userid string) (*RedisUser, error) {
	key := fmt.Sprintf("%s:%s", config.PreixForUser, userid)
	var errStr string
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()
	var redisUser *RedisUser
	var redisMap map[string]string
	out := GetTokenByUserid(ctx, key)
	select {
	case <-ctx.Done():
		errStr = "call get token by userId timeout"
	case rm := <-out:
		redisMap = rm
		mapstructure.Decode(redisMap, &redisUser)
		if redisUser.Nickname != "" && redisUser.Headimgurl != "" {
			return redisUser, nil
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}

	//没有就查mongodb
	var resultMongo user.User
	in3 := checkUserInMongoByUserid(userid)
	select {
	case <-ctx.Done():
		errStr = "query mongo timeout"
	case rm := <-in3:
		resultMongo = rm
		if resultMongo.Id == "" {
			errStr = "no user in mongo"
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}

	redisUser = &RedisUser{
		Id:         resultMongo.Id.Hex(),
		Unionid:    resultMongo.Unionid,
		MinAppid:   resultMongo.MinAppid,
		MinOpenid:  resultMongo.MinOpenid,
		Nickname:   resultMongo.Nickname,
		Headimgurl: resultMongo.Headimgurl,
		Sex:        resultMongo.Sex,
		City:       resultMongo.City,
		Province:   resultMongo.Province,
		Country:    resultMongo.Country,
		CreateTime: strconv.Itoa(int(resultMongo.CreateTime)),
	}
	in5 := UpdateUserToRedis(ctx, key, redisUser, strconv.Itoa(3600*24*15))
	select {
	case <-ctx.Done():
		errStr = "save redis timeout"
	case ok := <-in5:
		if ok != "OK" {
			errStr = ok
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}
	//fmt.Println("save redis successful .. ")
	return redisUser, nil
}
func checkUserInMongoByUserid(userid string) chan user.User {
	out := make(chan user.User)
	go func() {
		defer close(out)
		var u user.User
		u, err := user.FindById(userid, nil)
		if err != nil {
			out <- u
			return
		}
		out <- u
	}()
	return out
}

type WxUser struct {
	Signature     string `json:"signature"`
	RawData       string `json:"rawData"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

func SaveMinUser(ctx iris.Context) {
	userid := ctx.URLParam("userid")
	token := ctx.URLParam("token")

	var rawData, signature, encryptedData, iv string

	//fmt.Println(ctx.GetContentTypeRequested(),"-------")
	if strings.Contains(ctx.GetContentTypeRequested(), "json") {
		//application/json
		c := &WxUser{}
		if err := ctx.ReadJSON(c); err != nil {
			ctx.JSONP(iris.Map{"status": 401, "msg": "pls send application/json"})
			return
		}
		rawData = c.RawData
		signature = c.Signature
		encryptedData = c.EncryptedData
		iv = c.Iv
	} else {
		//application/x-www-form-urlencoded
		rawData = ctx.FormValue("rawData")
		signature = ctx.FormValue("signature")
		encryptedData = ctx.FormValue("encryptedData")
		iv = ctx.FormValue("iv")
	}
	//check 参数
	if userid == "" || token == "" || signature == "" || rawData == "" || encryptedData == "" || iv == "" {
		ctx.JSONP(iris.Map{"status": 401, "msg": "缺少必传参数"})
		return
	}
	//check redis token by user
	var redisUser *RedisUser
	key := fmt.Sprintf("%s:%s", config.PreixForUser, userid)
	if redisUserOld, err := CheckToken(key, token); err != nil {
		ctx.JSONP(iris.Map{"status": 401, "msg": err.Error()})
		return
	} else {
		redisUser = redisUserOld
	}

	//check signature
	sign := utils.Sha1s(rawData + redisUser.SessionKey)

	if sign != signature {
		fmt.Println(redisUser.Id, redisUser.MinOpenid, signature, "!=", sign)
		emsg := fmt.Sprintf("signature not equal (req/res): %s!=%s", signature, sign)
		ctx.JSONP(iris.Map{"status": 401, "msg": emsg})
		return
	}

	//check encrypteData
	str, err := minauth.DecryptData(encryptedData, redisUser.SessionKey, iv)
	if err != nil {
		ctx.JSONP(iris.Map{"status": 401, "msg": err.Error()})
		return
	}
	userinfo := &WxUserinfo{}
	models.UnmarshalNew([]byte(str), userinfo)

	//check appid all same
	waterMark := userinfo.Watermark
	if redisUser.MinOpenid != userinfo.OpenId && redisUser.MinAppid != waterMark.Appid {
		ctx.JSONP(iris.Map{"status": 401, "msg": "not same user"})
		return
	}

	//update user in mongo
	//update user in redis and token
	redisUser, err = UpdateUserInMongoAndRedis(key, userid, redisUser, userinfo)
	token = redisUser.Token

	redisUser.MinAppid = ""
	redisUser.MinOpenid = ""
	redisUser.Unionid = ""
	redisUser.SessionKey = ""
	redisUser.CreateTime = ""
	redisUser.UpdateTime = ""
	redisUser.Token = ""

	r, _ := models.Marshal(redisUser)

	var (
		result            *simplejson.Json
		isCallSaveMinUser bool
	)
	if redisUser.Nickname == "" && redisUser.Headimgurl == "" {
		isCallSaveMinUser = false
	} else {
		isCallSaveMinUser = true
		result, _ = simplejson.NewJson([]byte(r))
	}
	ctx.JSONP(iris.Map{"status": 200, "isCallSaveMinUser": isCallSaveMinUser,
		"data": result, "userid": userid, "token": token})

}

func CheckToken(key, token string) (*RedisUser, error) {
	var errStr string
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()
	var redisUser *RedisUser
	var redisMap map[string]string
	out := GetTokenByUserid(ctx, key)
	select {
	case <-ctx.Done():
		errStr = "call get token by userId timeout"
	case rm := <-out:
		redisMap = rm
		mapstructure.Decode(redisMap, &redisUser)
		if redisUser.Token != token {
			errStr = "invalid token"
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}
	return redisUser, nil
}

func GetTokenByUserid(ctx context.Context, key string) chan map[string]string {
	out := make(chan map[string]string)
	go func() {
		defer close(out)
		var hashobj map[string]string
		hashobj, err := lib.JudgeHgetall(key)
		if err != nil {
			out <- hashobj
			return
		}
		out <- hashobj
	}()
	return out
}

func UpdateUserInMongoAndRedis(key, userid string, redisUser *RedisUser, userinfo *WxUserinfo) (*RedisUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	in := UpdateUserToMongo(ctx, userid, bson.M{
		"minappid":  redisUser.MinAppid,
		"minopenid": redisUser.MinOpenid,
		"unionid":   redisUser.Unionid,

		"nickname":   userinfo.NickName,
		"headimgurl": userinfo.AvatarUrl,
		"province":   userinfo.Province,
		"country":    userinfo.Country,
		"city":       userinfo.City,
		"sex":        strconv.Itoa(userinfo.Gender),
		"createTime": redisUser.CreateTime,
		"updateTime": utils.GetCurrentTime(13),
	})
	var errStr string
	select {
	case <-ctx.Done():
		errStr = "query redis timeout"
	case id := <-in:
		if id == "" {
			errStr = "update mongo error"
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}

	//save to redis
	userStr, _ := models.Marshal(userinfo)
	tokenStr, _ := utils.GenJwtToken(userid, userStr, 15000)
	token := utils.Md5(tokenStr)

	redisUser.Nickname = userinfo.NickName
	redisUser.Headimgurl = userinfo.AvatarUrl
	redisUser.Sex = strconv.Itoa(userinfo.Gender)
	redisUser.City = userinfo.City
	redisUser.Province = userinfo.Province
	redisUser.Country = userinfo.Country
	redisUser.Token = token

	in2 := UpdateUserToRedis(ctx, key, redisUser, strconv.Itoa(3600*24*15))
	select {
	case <-ctx.Done():
		errStr = "query redis timeout"
	case reply := <-in2:
		if reply != "OK" {
			errStr = "update redis error"
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}

	return redisUser, nil
}

func UpdateUserToMongo(ctx context.Context, id string, modify bson.M) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		err := user.UpdateById(id, modify)
		if err != nil {
			out <- ""
			return
		}
		out <- id
	}()
	return out
}

func UpdateUserToRedis(ctx context.Context, key string, v *RedisUser, ttl string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		m := v.StructToMap()
		if reply, err := lib.JudgeHmset(key, m, ttl); err != nil {
			r := fmt.Sprintf("save user and openid failed %s", reply)
			out <- r
			return
		}
		out <- "OK"
	}()
	return out
}

func Minlogin(ctx iris.Context) {
	code := ctx.URLParam("code")
	appid := ctx.URLParam("appid")

	var secret string
	allmap.Appids.Mu.RLock()
	if st, ok := allmap.Appids.Ids[appid]; !ok {
		ctx.JSONP(iris.Map{"status": 402, "msg": "invalid appid"})
		allmap.Appids.Mu.RUnlock()
		return
	} else {
		secret = st
	}
	allmap.Appids.Mu.RUnlock()

	body := &WxauthRequestBody{
		appid:  appid,
		secret: secret,
		code:   code,
	}
	redisUser, err := body.GetMinappAtuh()
	//fmt.Printf("appid:%s,code:%s",appid,code)
	if err != nil {
		ctx.JSONP(iris.Map{"status": 401, "msg": err.Error()})
		return
	}
	userid := redisUser.Id
	token := redisUser.Token

	redisUser.MinAppid = ""
	redisUser.MinOpenid = ""
	redisUser.Unionid = ""
	redisUser.SessionKey = ""
	redisUser.CreateTime = ""
	redisUser.UpdateTime = ""
	redisUser.Token = ""

	r, _ := models.Marshal(redisUser)

	var (
		result            *simplejson.Json
		isCallSaveMinUser bool
	)
	if redisUser.Nickname == "" && redisUser.Headimgurl == "" {
		isCallSaveMinUser = true
	} else {
		isCallSaveMinUser = false
		result, _ = simplejson.NewJson([]byte(r))
	}
	ctx.JSONP(iris.Map{"status": 200, "isCallSaveMinUser": isCallSaveMinUser,
		"data": result, "userid": userid, "token": token})
}

func (wb *WxauthRequestBody) GetMinappAtuh() (RedisUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()
	//微信code换openid
	url := fmt.Sprintf("%s/sns/jscode2session?appid=%s&secret=%s"+
		"&js_code=%s&grant_type=authorization_code", config.WxAuthHost, wb.appid, wb.secret, wb.code)
	in := minauth.GetMinappOpenid(ctx, url)
	//fmt.Println("--url--",url)
	var (
		errStr     string
		userKey    string
		userMapKey string
	)
	var resultWX minauth.MinAppOpenid
	var redisUser RedisUser
	select {
	case <-ctx.Done():
		errStr = "call wx code change openid timeout"
	case rm := <-in:
		resultWX = *rm
		//fmt.Println("--resultWX--", resultWX)
		if resultWX.Errcode != 0 && resultWX.Errmsg != "" {
			errStr = resultWX.Errmsg
		} else if resultWX.Openid == "" || resultWX.SessionKey == "" {
			errStr = "wx auth failed"
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}
	//通过appid和openid查询redis是否有这个用户
	userMapKey = fmt.Sprintf("%s:%s:%s", config.MapOpenidUserid, wb.appid, resultWX.Openid)
	in2 := wb.checkOpenidUserid(ctx, userMapKey)
	var redisMap map[string]string
	select {
	case <-ctx.Done():
		errStr = "query redis timeout"
	case rm := <-in2:
		redisMap = rm
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}
	//把redis的map转为实际的struct
	var resultRedis MapOpenidUserid
	mapstructure.Decode(redisMap, &resultRedis)

	if resultRedis.Userid != "" {
		userKey = fmt.Sprintf("%s:%s", config.PreixForUser, resultRedis.Userid)
		redisMap, err := lib.JudgeHgetall(userKey)
		if err != nil {
			return redisUser, nil
		}
		mapstructure.Decode(redisMap, &redisUser)
		if redisUser.Token == "" {
			goto QueryMongo
		} else if redisUser.SessionKey != resultWX.SessionKey {
			//另起gorutine把session_key更进去
			go func() {
				tmpUser := RedisUser{
					SessionKey: resultWX.SessionKey,
				}
				fmt.Printf("session_key expired, old:%s, new:%s\n", redisUser.SessionKey, resultWX.SessionKey)
				lib.JudgeHmset(userKey, tmpUser.StructToMap(), strconv.Itoa(3600*24*15))
			}()
		}
		//fmt.Println("redis have data ...", redisUser)
		return redisUser, nil
	}
	//fmt.Println("redis中没有")
QueryMongo:
	var mongoid string

	//没有就查mongodb
	var resultMongo *user.User
	in3 := wb.checkUserInMongo(resultWX.Openid)
	select {
	case <-ctx.Done():
		errStr = "query mongo timeout"
	case rm := <-in3:
		resultMongo = rm
		if resultMongo != nil {
			mongoid = resultMongo.Id.Hex()
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}

	if mongoid == "" {
		//如果mongodb也没有，就存到mongodb中
		resultMongo = &user.User{
			MinOpenid:  resultWX.Openid,
			MinAppid:   wb.appid,
			Unionid:    resultWX.Unionid,
			CreateTime: utils.GetCurrentTime(13),
		}
		in4 := wb.SaveUserToMongo(resultMongo)
		select {
		case <-ctx.Done():
			errStr = "save mongo timeout"
		case r := <-in4:
			mongoid = r
			//fmt.Println("不在mongo中，但存成功", mongoid)
		}
		if errStr != "" {
			return redisUser, errors.New(errStr)
		}
	}

	//fmt.Println("mongo中有数据", mongoid)

	//然后把用户存入redis及 appid:openid和userid的映射关系
	resultRedis = MapOpenidUserid{
		Userid:  mongoid,
		Unionid: resultMongo.Unionid,
	}

	userInMongoStr, _ := models.Marshal(resultMongo)
	tokenStr, _ := utils.GenJwtToken(mongoid, userInMongoStr, 15000)
	token := utils.Md5(tokenStr)
	redisUser = RedisUser{
		Id:         mongoid,
		Unionid:    resultMongo.Unionid,
		MinAppid:   resultMongo.MinAppid,
		MinOpenid:  resultMongo.MinOpenid,
		Nickname:   resultMongo.Nickname,
		Headimgurl: resultMongo.Headimgurl,
		Sex:        resultMongo.Sex,
		City:       resultMongo.City,
		Province:   resultMongo.Province,
		Country:    resultMongo.Country,
		CreateTime: strconv.Itoa(int(resultMongo.CreateTime)),
		UpdateTime: strconv.Itoa(int(resultMongo.UpdateTime)),
		SessionKey: resultWX.SessionKey,
		Token:      token,
	}
	//fmt.Println("=======", resultRedis)
	//fmt.Println("=======", redisUser)
	userKey = fmt.Sprintf("%s:%s", config.PreixForUser, resultRedis.Userid)
	in5 := wb.SaveUserToHashAndHash(ctx, userKey, userMapKey, redisUser, resultRedis)
	select {
	case <-ctx.Done():
		errStr = "save redis timeout"
	case ok := <-in5:
		if ok != "OK" {
			errStr = ok
		}
	}
	if errStr != "" {
		return redisUser, errors.New(errStr)
	}
	//fmt.Println("save redis successful .. ")
	return redisUser, nil
}

func (wb *WxauthRequestBody) checkOpenidUserid(ctx context.Context, key string) chan map[string]string {
	out := make(chan map[string]string)
	go func() {
		defer close(out)
		var openidUserid map[string]string
		openidUserid, err := lib.JudgeHgetall(key)
		if err != nil {
			out <- openidUserid
			return
		}
		out <- openidUserid
	}()

	return out
}

func (wb *WxauthRequestBody) checkUserInMongo(openid string) chan *user.User {
	out := make(chan *user.User)
	go func() {
		defer close(out)
		var u *user.User
		u, err := user.FindByAppidOpenid(wb.appid, openid)
		if err != nil {
			out <- u
			return
		}
		out <- u
	}()
	return out
}

func (wb *WxauthRequestBody) SaveUserToMongo(userinfo *user.User) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		id, err := user.SaveUser(userinfo)
		if err != nil {
			out <- ""
			return
		}
		out <- id
	}()
	return out
}

func (wb *WxauthRequestBody) SaveUserToHashAndHash(ctx context.Context, key,
	key2 string, v RedisUser, v2 MapOpenidUserid) chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		m := v.StructToMap()
		m2 := v2.StructToMap()

		reply, reply2 := lib.JudgeHmsetAndHmset(key, m, key2, m2, strconv.Itoa(3600*24*15))

		if reply != "OK" && reply2 != "OK" {
			r := fmt.Sprintf("save user and openid failed %s", reply)
			out <- r
			return
		}
		out <- "OK"
	}()
	return out
}
