package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestData 请求数据结构体
type RequestData struct {
	ReqBody      string `json:"reqBody"`
	KeyEncrypted string `json:"keyEncrypted"`
}

// ResponseBodyS 响应体结构体
type ResponseBodyS struct {
	Result string `json:"result"`
}

func mainaas() {
	http.HandleFunc("/process", processData)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func processData(w http.ResponseWriter, r *http.Request) {
	// 读取并解析请求体
	reqBodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	reqData := RequestData{}
	if err := json.Unmarshal(reqBodyByte, &reqData); err != nil {
		panic(err)
	}

	// 服务端使用私钥解密 AES 密钥
	keyEncryptedByte, err := base64.StdEncoding.DecodeString(reqData.KeyEncrypted)
	if err != nil {
		panic(err)
	}
	key, err := rsaDecrypt(keyEncryptedByte)
	if err != nil {
		panic(err)
	}

	// 服务端使用 AES 解密请求体
	reqBodyEncryptedByte, err := base64.StdEncoding.DecodeString(reqData.ReqBody)
	if err != nil {
		panic(err)
	}
	reqBodyByte, err = decryptAES(key, reqBodyEncryptedByte)
	if err != nil {
		panic(err)
	}

	// 服务端处理请求体，这里假设直接返回请求体
	respData := ResponseBodyS{
		Result: string(reqBodyByte),
	}
	respDataByte, err := json.Marshal(respData)
	if err != nil {
		panic(err)
	}
	_, _ = w.Write(respDataByte)
}

// rsaDecrypt 使用 RSA 算法对数据进行解密
func rsaDecrypt(data []byte) ([]byte, error) {
	privateKey, err := readRSAPrivateKeyFromFile("private.key")
	if err != nil {
		return nil, err
	}
	sum, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		return nil, err
	}
	return sum, nil
}

// decryptAES 使用 AES 算法对数据进行解密
func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return PKCS7Unpadding(ciphertext), nil
}

// PKCS7Unpadding 对数据进行 PKCS7 反填充
func PKCS7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func generateRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// 生成 RSA 密钥对并保存到文件
func generateAndSaveRSAKey() error {
	privatekey, err := generateRSAKey()
	if err != nil {
		return err
	}
	if err := saveRSAPrivateKeyToFile(privatekey, "private.key"); err != nil {
		return err
	}
	publickeyBytes, err := x509.MarshalPKIXPublicKey(&privatekey.PublicKey)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("public.key", publickeyBytes, 0644); err != nil {
		return err
	}
	return nil
}

// 从文件中读取 RSA 私钥
func readRSAPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(privateKeyBytes)
}

// 将私钥保存到文件
func saveRSAPrivateKeyToFile(key *rsa.PrivateKey, filename string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	return ioutil.WriteFile(filename, privateKeyBytes, 0600)
}

// 从文件中读取 RSA 公钥
func readRSAPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not a public key")
	}
	return publicKey, nil
}

// 使用 RSA 算法对数据进行加密
/*func rsaEncrypt(pubkey *rsa.PublicKey, data []byte) ([]byte, error) {
	sum := sha256.Sum256(data)
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubkey, sum[:], nil)
}*/

// 使用 AES 算法对数据进行加密
/*func encryptAES(key, data []byte) ([]byte, error) {
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
}*/

// 对数据进行 PKCS7 填充
/*func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}*/

/*func a() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	// Create public key from private key
	publicKey := privateKey.PublicKey

	// Save private key to file
	privateFile, err := os.Create("private.key")
	defer privateFile.Close()
	if err != nil {
		panic(err)
	}
	privateEncoder := gob.NewEncoder(privateFile)
	err = privateEncoder.Encode(privateKey)
	if err != nil {
		panic(err)
	}

	// Save public key to file
	publicFile, err := os.Create("public.key")
	defer publicFile.Close()
	if err != nil {
		panic(err)
	}
	publicAsn1, err := asn1.Marshal(publicKey)
	if err != nil {
		panic(err)
	}
	publicEncoder := pem.Encode(publicFile, &pem.Block{Type: "PUBLIC KEY", Bytes: publicAsn1})
	if err != nil {
		panic(err)
	}
}*/

func main() {
	/*	// Generate RSA key pair
		privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			panic(err)
		}

		// Create public key from private key
		publicKey := privateKey.PublicKey

		// Save private key to file
		privateFile, err := os.Create("private.key")
		defer privateFile.Close()
		if err != nil {
			panic(err)
		}
		privateEncoder := gob.NewEncoder(privateFile)
		err = privateEncoder.Encode(privateKey)
		if err != nil {
			panic(err)
		}

		// Save public key to file
		publicFile, err := os.Create("public.key")
		defer publicFile.Close()
		if err != nil {
			panic(err)
		}
		publicAsn1, err := asn1.Marshal(publicKey)
		if err != nil {
			panic(err)
		}
		_ = pem.Encode(publicFile, &pem.Block{Type: "PUBLIC KEY", Bytes: publicAsn1})
		if err != nil {
			panic(err)
		}*/
	generateAndSaveRSAKey()
}
