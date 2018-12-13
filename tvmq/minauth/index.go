package minauth

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/models"
)

type MinAppOpenid struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid,omitempty"`
}

func GetMinappOpenid(ctx context.Context, url string) chan *MinAppOpenid {
	out := make(chan *MinAppOpenid)
	go func() {
		defer close(out)

		//xr := &MinAppOpenid{
		//	Errcode: 0,
		//	Openid:     "o8YpOxAjDD1dQqPxSpTFySa2U75s" + strconv.Itoa(rand.Int()),
		//	//Openid:     "o8YpOxAjDD1dQqPxSpTFySa2U75s",
		//	Unionid:    "oDFXz0J5eEXbg1jg_G2uUrc-uLnY",
		//	SessionKey: "kZzq4Dgqg4Yx+W+WVlXTkA==",
		//}
		//out <- xr
		//return

		s, err := backend.SendGet(url)
		r := &MinAppOpenid{}
		if err != nil {
			fmt.Println(err)
			out <- r
			return
		}
		models.UnmarshalNew(s, r)
		out <- r
	}()
	return out
}

func AesCBCDncrypt(encryptData, key, iv []byte) (string, error) {
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(encryptData))
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.CryptBlocks(decrypted, encryptData)

	return string(decrypted), nil
}

func DecryptData(encryptedData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	iKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	iIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	dnData, err := AesCBCDncrypt(data, iKey, iIv)
	if err != nil {
		return "", err
	}
	return dnData, nil
}
