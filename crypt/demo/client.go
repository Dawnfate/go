package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// RequestBody 请求体结构体
type RequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ResponseBody 响应体结构体
type ResponseBody struct {
	Result string `json:"result"`
}

func main() {
	// 客户端生成 AES 密钥
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	// 客户端加密请求体
	reqBody := RequestBody{
		Name: "John",
		Age:  28,
	}
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	reqBodyEncrypted, err := encryptAES(key, reqBodyByte)
	if err != nil {
		panic(err)
	}

	// 客户端使用服务端公钥加密 AES 密钥
	pubKey := getPublicKey("./crypt/demo/public.key")

	keyEncrypted, err := rsaEncrypt(pubKey, key)
	if err != nil {
		panic(err)
	}

	// 构造 HTTP 请求体
	reqData := map[string]interface{}{
		"reqBody":      base64.StdEncoding.EncodeToString(reqBodyEncrypted),
		"keyEncrypted": base64.StdEncoding.EncodeToString(keyEncrypted),
	}
	reqDataByte, err := json.Marshal(reqData)
	if err != nil {
		panic(err)
	}

	// 发送 HTTP 请求
	resp, err := http.Post("http://localhost:8080/encrypt", "application/json", bytes.NewReader(reqDataByte))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 解析响应体
	respBodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respData := ResponseBody{}
	if err := json.Unmarshal(respBodyByte, &respData); err != nil {
		panic(err)
	}
	fmt.Println(respData.Result)
}

func getPublicKey(path string) *rsa.PublicKey {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey
}

// encryptAES 使用 AES 算法对数据进行加密
func encryptAES(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	dataPKCS7Padded := PKCS7Padding(data, blockSize)
	ciphertext := make([]byte, block.BlockSize()+len(dataPKCS7Padded))
	iv := ciphertext[:block.BlockSize()]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[block.BlockSize():], dataPKCS7Padded)
	return ciphertext, nil
}

// rsaEncrypt 使用 RSA 算法对数据进行加密
func rsaEncrypt(pubKey *rsa.PublicKey, data []byte) ([]byte, error) {
	sum := sha256.Sum256(data)
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, sum[:], nil)
}

//todo getPublicKey 获取服务端公钥
/*func getPublicKey() (*rsa.PublicKey, error) {
	// 实际中需要替换成从服务端获取公钥的代码
	return rsa.GenerateKey(rand.Reader, 2048).PublicKey, nil
}*/

// PKCS7Padding 对数据进行 PKCS7 填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
