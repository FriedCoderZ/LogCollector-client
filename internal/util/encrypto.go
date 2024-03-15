package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

// RSA加密函数
func EncryptedAESKey(aesKey []byte, publicKeyPath string) (string, error) {
	// 读取RSA公钥文件
	publicKeyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return "", nil
	}

	// 解析RSA公钥
	block, _ := pem.Decode(publicKeyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return "", fmt.Errorf("failed to decode RSA public key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", nil
	}

	// 使用RSA公钥加密AES密钥
	encryptedKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, aesKey)
	if err != nil {
		return "", nil
	}

	// 将加密后的AES密钥进行Base64编码
	encryptedKeyBase64 := base64.StdEncoding.EncodeToString(encryptedKey)
	return encryptedKeyBase64, nil
}

func GenerateAESKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext, nil
}
