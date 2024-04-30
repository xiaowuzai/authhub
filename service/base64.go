package service

import (
	"encoding/base64"
)

func StrToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func StrFromBase64(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(b), err
}
