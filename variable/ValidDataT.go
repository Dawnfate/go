package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func FilterInvalidData(data string) string {
	reg := regexp.MustCompile("^[\u4e00-\u9fa5A-Za-z0-9]$")
	result := ""
	for _, v := range data {
		if reg.MatchString(string(v)) {
			result = result + string(v)
		}
	}
	return result
}

type Head struct {
	APPID     string `json:"APP_ID"`
	TIMESTAMP string `json:"TIMESTAMP"`
	NONCE     string `json:"NONCE"`
	TOKEN     string `json:"TOKEN"`
}

type VerifyHead struct {
	Head
	AppSecret string
}

func aba(head VerifyHead) string {
	str := "APP_ID" + head.APPID + "TIMESTAMP" + head.TIMESTAMP + "NONCE" + head.NONCE + head.AppSecret
	return str
}

//返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	ab := time.Now().Unix()
	fmt.Println(ab)
	a := strconv.FormatInt(ab, 10)
	verifyHead := VerifyHead{
		Head{
			APPID:     "yaditrans",
			TIMESTAMP: a,
			NONCE:     "1231231232",
			TOKEN:     "",
		},
		"",
	}
	fmt.Println(a)
	fmt.Println(GetMD5Encode(aba(verifyHead)))
}
