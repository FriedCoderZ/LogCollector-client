package collect

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

// RSA加密函数
func rsaEncrypt(plainText []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
}

func main() {
	// 读取RSA公钥文件
	publicKeyFile := "publicKey.pem"
	publicKeyData, err := os.ReadFile(publicKeyFile)
	if err != nil {
		log.Fatal("Failed to read public key file:", err)
		return
	}

	// 解析RSA公钥
	block, _ := pem.Decode(publicKeyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("Failed to decode RSA public key")
		return
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Fatal("Failed to parse RSA public key:", err)
		return
	}

	// 待加密的AES密钥
	aesKey := []byte("MySecretAESKey")

	// 使用RSA公钥加密AES密钥
	encryptedKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, aesKey)
	if err != nil {
		log.Fatal("Failed to encrypt AES key with RSA:", err)
		return
	}

	// 将加密后的AES密钥进行Base64编码
	encryptedKeyBase64 := base64.StdEncoding.EncodeToString(encryptedKey)
	fmt.Println("Encrypted AES key:", encryptedKeyBase64)
}
