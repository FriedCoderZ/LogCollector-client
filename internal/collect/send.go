package collect

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/FriedCoderZ/LogCollector-client/internal/config"
	"github.com/FriedCoderZ/LogCollector-client/internal/util"
)

func register() error {
	// 生成AES密钥
	aesKey, err := util.GenerateAESKey(128)
	if err != nil {
		return err
	}
	// 加密AES密钥
	config := config.GetConfig()
	publicKeyPath := config.Crypto.RSAPublicKeyPath
	encryptedKey, err := util.EncryptedAESKey(aesKey, publicKeyPath)
	if err != nil {
		return err
	}

	// 发送加密后的AES密钥到服务器
	err = sendEncryptedKey(encryptedKey)
	if err != nil {
		return err
	}

	fmt.Println("AES密钥已经发送至服务器")
	return nil
}

func sendEncryptedKey(encryptedKey string) error {
	url := "http://127.0.0.1:8080" // 服务器地址
	payload := strings.NewReader(encryptedKey)

	resp, err := http.Post(url, "application/octet-stream", payload)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
