package msgtype

import (
	"code.tvmining.com/tvplay/tvmq/logs"
	"code.tvmining.com/tvplay/tvmq/logs/trace"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/utils"
	"github.com/bitly/go-simplejson"
)

func CheckMessageType(chatFlag, message, from string) (string, string) {
	var logger logs.Logger
	var traceId string
	var res string
	if from == "" {
		from = utils.GetIntranetIp()
	}

	body, _ := simplejson.NewJson([]byte(message))
	rbMessage, _ := body.Get("Message").String()

	ti := body.Get("TraceInfo")
	if ti != nil {
		traceId, _ = ti.Get("traceId").String()
		logger = &trace.TraceInfo{
			TraceId:    traceId,
			RecvByNode: utils.GetCurrentTime(13),
		}
		logs.TraceChan <- logger
	}

	switch chatFlag {
	case "chat":
		d := models.ResComment{}
		models.Unmarshal(rbMessage, &d)
		//d.Unmarshal(message)

		reply := models.ResChatBody{
			Status:  200,
			Content: d,
			From:    from,
		}
		res, _ = models.Marshal(reply)

	case "gift":
		d := models.ResGift{}
		models.Unmarshal(rbMessage, &d)

		//d.Unmarshal(message)
		reply := models.ResGiftBody{
			Status:  200,
			Content: d,
			From:    from,
		}
		res, _ = models.Marshal(reply)
	case "newAuthUser":
		d := models.ResComment{}
		models.Unmarshal(rbMessage, &d)

		reply := models.ResChatBody{
			Status:  200,
			Content: d,
			From:    from,
		}
		res, _ = models.Marshal(reply)
	default:

		//d := models.ResFormanBody{}
		//models.Unmarshal(message, &d)
		////d.Unmarshal(message)
		//content,_ := models.Marshal(d.Content)
		//reply := models.ResFormanBody{
		//	Status:  200,
		//	Content: content,
		//	From:    from,
		//}
		//res, _ = models.Marshal(reply)

		res = message

	}
	return res, traceId
}
