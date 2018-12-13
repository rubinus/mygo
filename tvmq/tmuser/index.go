package tmuser

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	userTokenVersion1 = "1"
)

type UserPayload struct {
	UserId  string `json:"u,omitempty"`
	Expired uint32 `json:"e,omitempty"`
}

func (it *UserPayload) IsExpired() bool {
	return it.Expired <= uint32(time.Now().Unix())
}

type UserPayloadInterface interface {
	IsExpired() bool
}

func NewToken(user UserPayloadInterface, secret_key string) string {

	var (
		bs, _   = json.Marshal(user)
		payload = strings.TrimRight(base64.URLEncoding.EncodeToString(bs), "=")
	)

	return fmt.Sprintf("%s.%s.%s",
		userTokenVersion1,
		payload,
		TokenSign(userTokenVersion1, payload, secret_key),
	)
}

func TokenValid(token, secret_key string, payload UserPayloadInterface) error {

	var (
		n1 = strings.IndexByte(token, '.')
		n2 = strings.LastIndexByte(token, '.')
	)

	if 0 < n1 && n1 < n2 && n2+2 < len(token) {

		switch token[:n1] {

		case userTokenVersion1:

			bs, err := base64.URLEncoding.DecodeString(base64fix(token[n1+1 : n2]))
			if err != nil {
				return errors.New("invalid payload " + err.Error())
			}

			if err = json.Unmarshal(bs, payload); err != nil {
				return errors.New("invalid payload/json " + err.Error())
			}

			if payload.IsExpired() {
				return errors.New("tmt expired")
			}

			if TokenSign(userTokenVersion1, token[n1+1:n2], secret_key) != token[n2+1:] {
				return errors.New("sign denied")
			}

			return nil
		}
	}

	return errors.New("invalid token")
}

func TokenSign(version, payload, secret_key string) string {
	hs := sha256.Sum256([]byte(payload + secret_key))
	return strings.TrimRight(base64.URLEncoding.EncodeToString(hs[:16]), "=")
}

func base64fix(s string) string {
	if n := len(s) % 4; n > 0 {
		s += strings.Repeat("=", 4-n)
	}
	return s
}
