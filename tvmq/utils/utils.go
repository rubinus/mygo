package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"code.tvmining.com/tvplay/tvmq/tmuser"
	"github.com/dgrijalva/jwt-go"
	"github.com/json-iterator/go"

	"code.tvmining.com/tvplay/tis/tiapi"
)

func Marshal(res interface{}) (string, error) {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	s, err := jsonIterator.Marshal(res)
	return string(s), err
}

func GetCurrentTime(f int) int64 {
	var result int64
	switch f {
	case 10:
		result = time.Now().Unix()
	case 13:
		result = time.Now().UnixNano() / 1e6
	case 19:
		result = time.Now().UnixNano()
	}
	return result
}

var (
	ip_cache = ""
)

func GetIntranetIp() string {
	if ip_cache != "" {
		return ip_cache
	}

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return "127.0.0.1"
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ip_cache = ipnet.IP.String()
						break
					}
				}
			}
		}
	}

	return ip_cache
}

func ParseDns(strDns string) bool {
	ns, err := net.LookupHost(strDns)
	if err != nil {
		fmt.Printf("error: %v, failed to parse %v\n", err, strDns)
		return false
	}

	if len(ns) <= 0 {
		return false
	}

	return true
}

func GenJwtToken(key, body string, expire int64) (string, error) {
	mySigningKey := []byte(key)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expire,
		Issuer:    body,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func TokenValid(token, key string) (string, string, string, error) {
	var item tiapi.UserTokenPayload
	if err := tmuser.TokenValid(token, key, &item); err != nil {
		return "", "", "", err
	} else {
		return item.UserPayload.UserId, item.NickName, item.AvatarUrl, nil
	}
}

func Md5(body string) string {
	md5 := md5.New()
	md5.Write([]byte(body))
	return hex.EncodeToString(md5.Sum(nil))
}

func Sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}
