package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

var StartKafkaTestData, StartKafkaConsumer, StartNsqConsumer, StartRpcServe, UseRedisCluster,
	UseMongodbCluster, UseFilterWord, UseSendMsgLimit, SendMsgLimitSecond, WebServerPort,
	SelfServePort, UsePreAbsPath, RecentMsgCount int

var NsqHostProt, SocketHost, PointHost, FilterHost, DefaultRedisHost, DefaultMongoHost string

var DefaultRedisPort, MongoDBPort, TraceLog int

var EnvHost bool

var Minappid, Minsecret, TokenSecretKey string

var MongoDatabase, MongoUsername, MongoPassword string

var RedisClusterIP, MongodbClusterIP, KafkaHostPort []string

func init() {
	var env string
	flag.StringVar(&env, "env", "dev", "start dev or qa or pro env config")
	flag.Parse()
	initConfig("./config/" + env + ".json")
}

func initConfig(cf string) {
	viper.SetConfigFile(cf)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("env"), "is started ...")

	//使用docker container用81 ，外部请求连接端口lb. 本地用8081测试
	WebServerPort = viper.GetInt("WebServerPort")

	//本程序自己的端口，勿动
	SelfServePort = viper.GetInt("SelfServePort")

	////true用本地，false用wss，socket连接时使用的ws://或wss://，正式用false
	EnvHost = viper.GetBool("EnvHost")

	//nsq地址
	NsqHostProt = viper.GetString("NsqHostProt")

	//redis地址
	RedisClusterIP = viper.GetStringSlice("RedisClusterIP")

	//mongodb地址
	MongodbClusterIP = viper.GetStringSlice("MongodbClusterIP")

	//kafka地址
	KafkaHostPort = viper.GetStringSlice("KafkaHostPort")

	//5秒一次向kafka发送测试数据 topic 是 socket_face
	StartKafkaTestData = viper.GetInt("StartKafkaTestData")

	//启动kafka的消费端
	StartKafkaConsumer = viper.GetInt("StartKafkaConsumer")

	//启动nsq的消费端
	StartNsqConsumer = viper.GetInt("StartNsqConsumer")

	//启动rpc服务端
	StartRpcServe = viper.GetInt("StartRpcServe")

	//1表示使用redis cluster
	UseRedisCluster = viper.GetInt("UseRedisCluster")

	//1表示使用mongodb shard cluster
	UseMongodbCluster = viper.GetInt("UseMongodbCluster")

	//1表示启用敏感词过滤功能，0表示不向过滤组件发请求
	UseFilterWord = viper.GetInt("UseFilterWord")

	//1表示启用发送消息限制，默认3秒内不能再次发消息
	UseSendMsgLimit = viper.GetInt("UseSendMsgLimit")

	//3秒内不能重发，和UseSendMsgLimit一起使用
	SendMsgLimitSecond = viper.GetInt("SendMsgLimitSecond")

	//1表示使用abs path
	UsePreAbsPath = viper.GetInt("UsePreAbsPath")

	//1表示记录追踪日志
	TraceLog = viper.GetInt("TraceLog")

	SocketHost = viper.GetString("SocketHost")
	PointHost = viper.GetString("PointHost")
	FilterHost = viper.GetString("FilterHost")
	Minappid = viper.GetString("Minappid")
	Minsecret = viper.GetString("Minsecret")

	TokenSecretKey = viper.GetString("TokenSecretKey")

	DefaultRedisHost = viper.GetString("redis.DefaultRedisHost")
	DefaultRedisPort = viper.GetInt("redis.DefaultRedisPort")

	DefaultMongoHost = viper.GetString("mongodb.DefaultMongoHost")
	MongoDatabase = viper.GetString("mongodb.MongoDatabase")
	MongoUsername = viper.GetString("mongodb.MongoUsername")
	MongoPassword = viper.GetString("mongodb.MongoPassword")
	MongoDBPort = viper.GetInt("mongodb.MongoDBPort")

	//最近消息总数
	RecentMsgCount = viper.GetInt("RecentMsgCount")

	rhost := os.Getenv("REDIS_HOST")
	mhost := os.Getenv("MONGO_HOST")
	nhost := os.Getenv("NSQ_HOST")
	khost := os.Getenv("KAFKA_HOST")
	if rhost != "" {
		DefaultRedisHost = rhost
	}
	if mhost != "" {
		MongodbClusterIP = []string{mhost}
	}
	if nhost != "" {
		NsqHostProt = nhost + ":4150"
	}
	if khost != "" {
		KafkaHostPort = []string{khost}
	}

	//fmt.Printf("rhost:%s, mhost:%s, nhost:%s, khost:%s", DefaultRedisHost, MongodbClusterIP, NsqHostProt,KafkaHostPort)
}

var (
	WxAuthHost      = "https://api.weixin.qq.com"
	MapOpenidUserid = "MAPOUID"
	RecentMsgKey    = "RECENT_MSG"

	JwtSecretKey    = "secret-key"
	DefalutTenantId = "100000"

	GoroutineTimeout = 5 * time.Second
	PointGetPath     = "/socialtv/api/v1/user/entry?user_id="
	PointPostPath    = "/socialtv/api/v1/gifts/gift-give"

	TopicEpg         = "epg"
	TopicLottery     = "lottery"
	TopicFace        = "face"
	TopicActivity    = "activity"
	TopicGift        = "nsgift"
	TopicNewAuthUser = "newAuthUser"
	TopicChat        = "chat"
	TopicForman      = "socialtv_im_queue11"
	TopicIMua        = "socialtv_im_ua"

	TraceInfo = "TraceInfo"

	RpcServerPort       = 3700
	RpcCallPath         = "/rpcsend"
	ChatRoomName        = "CHATROOM"
	SocEventHeartBeat   = "heartbeat"
	SocEventIP          = "ip"
	SocEventChat        = "chat"
	SocEventChatReply   = "chat_reply"
	SocEventAi          = "ai"
	SocEventAuth        = "auth"
	SocEventDiss        = "diss"
	SocEventGift        = "gift"
	SocEventGiftReply   = "gift_reply"
	SocEventActivity    = "activity"
	SocEventFace        = "face"
	SocEventEpg         = "epg"
	SocEventForman      = "forman"
	SocEventLottery     = "lottery"
	SocEventOnline      = "online" //在线人线
	SocEventNewAuthUser = "newAuthUser"

	PreixForUser      = "USER"
	ChatPreixForUser  = "SCU"
	LastSendPreix     = "LASTSEND"
	OnlineHostConnKey = "ALLHOSTS"

	CurrentActivity = "CurrentActivity"
)
