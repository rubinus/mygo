package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

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

func main() {
	var src = "L3RA2V84ZrBeiW1n3o5EuEbGPp0lAYDmSm20eaa+EjOR5QUp3cIS/l4frVugw8MtbuHQ8TYTJl3CNUDd7ZPJuOQWI54C/xKo0ixKS/PSZ2CQABNBZlQAN6/Fo6Ott3Oau18xXQ9car0yuSt9xnsjSXPMVLI7kIHNWtf3pf9XamN1lHMRYLn7f2osHu2uVE2LHRLvEAewqkAsnwMJ8PH3qCqyASpb4k0EieeDp/CrKbDyuF7dvnps5+vLObuEUIvkw6SqR/jjb9/ofTueEyRKyT0ST37/qedmFE5GzfdVx8y6ZBsHmJxokpLPTGpvi+bMR5BjC+zYnCCd0VC9j+aUl2gn5UHp3z6hj5yRbT4SaYr/VhVf5QBeMpIervLRn01x0+b2+Ab84RPQFaSngynPzvpQXuMGHI/SUCjid93ZCONpgResHvAoIhcvxHmYSCe6lmiM9C18Ew0h+f59siZpw4GUfrtqQm+I0w/Suv1dbb0/8N40wg8OqXmdmiWdX5uSJpKy20VnvkIT/fV3KyQ0q2H5xMOPh1ATux0OJA1JF3A="

	key := "kZzq4Dgqg4Yx+W+WVlXTkA=="
	iv := "jmRxicHfU4auUEMOwUkFsg=="
	s, _ := DecryptData(src, key, iv)

	fmt.Println(s)
}
