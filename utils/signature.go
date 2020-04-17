package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
)

func GenerateSignature(param url.Values, salt string) string {
	jsonStr, _ := json.Marshal(param)

	h := md5.New()
	h.Write(jsonStr)
	h.Write([]byte(salt))

	sign := hex.EncodeToString(h.Sum(nil))

	return sign
}

func VerifySignature(param url.Values, salt, key string) error {
	var err error

	signParam, ok := param[key]
	if !ok {
		err = errors.New("the sign is invaild.")
		return err
	}

	delete(param, key)
	sign := GenerateSignature(param, salt)

	if signParam[0] != sign {
		err = errors.New("request is illegal. the sign is not match.")
	}

	return err
}
