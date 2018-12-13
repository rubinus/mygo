package minauth

import (
	"context"
	"fmt"
	"testing"

	"code.tvmining.com/tvplay/tvmq/config"
)

func TestGetMinappOpenid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	appid := "wx8c86289016e25814"
	secret := "c1a208c9f23d44bf289d788ce64bf57e"
	code := "081JSai50cNvgK1gXqg50h49i50JSail"
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code)
	in := GetMinappOpenid(ctx, url)

	var errStr string
	select {
	case <-ctx.Done():
		errStr = "超时"
	case result := <-in:
		if result.Errcode != 0 && result.Errmsg != "" {
			errStr = result.Errmsg
		}
		fmt.Printf("%#v", result)
	}
	if errStr != "" {
		fmt.Println("::", errStr)
	}
}
