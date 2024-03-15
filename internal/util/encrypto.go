package util

import (
	"crypto/aes"
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
func EncryptRSA(publicKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, fmt.Errorf("decode public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

func EncryptAESKey(aesKey []byte, publicKeyPath string) (string, error) {
	// 读取RSA公钥文件
	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return "", err
	}

	encryptedKey, err := EncryptRSA(publicKey, aesKey)
	if err != nil {
		return "", err
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

func EncryptAES(plaintext string, key []byte) (string, error) {
	// 创建cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, len(plaintext))
	// 加密
	c.Encrypt(ciphertext, []byte(plaintext))
	// Base64编码返回
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertextBase64, nil
}

func DecryptAES(ciphertextBase64 string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(ciphertext))
	c.Decrypt(plaintext, ciphertext)

	plaintextStr := string(plaintext[:])
	return plaintextStr, nil
}
