package collect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/FriedCoderZ/LogCollector-client/internal/config"
	"github.com/FriedCoderZ/LogCollector-client/internal/database"
	"github.com/FriedCoderZ/LogCollector-client/internal/util"
)

// 在服务端注册登记采集端信息并商议AES密钥
func Register() error {
	// 获取配置信息
	config := config.GetConfig()

	// 生成AES密钥
	aesKey, err := util.GenerateAESKey(config.Crypto.AESLength)
	if err != nil {
		return err
	}

	// 加密AES密钥
	publicKeyPath := config.Crypto.RSAPublicKeyPath
	encryptedKey, err := util.EncryptAESKey(aesKey, publicKeyPath)
	if err != nil {
		return err
	}

	// 发送加密后的AES密钥到服务器
	url := config.Server.Address + "/collector"
	payload := strings.NewReader(encryptedKey)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var response struct {
			UUID string
		}
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			return err
		}
		database.CreateCollectInfo(response.UUID, aesKey)
		fmt.Printf("UUID:%s\n", response.UUID)
	} else {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func SendLogs(logs []map[string]string) error {
	if logs == nil {
		return fmt.Errorf("nil pointer dereference")
	}
	if len(logs) == 0 {
		return nil
	}
	config := config.GetConfig()
	serverAddress := config.Server.Address

	// 获取采集器信息
	collectorInfo, err := database.GetCollectorInfo()
	if err != nil {
		return err
	}

	// 提取采集器ID和AES密钥
	uuid := collectorInfo.UUID
	aesKey := collectorInfo.AESKey
	// 将ID和logs打包为一个json并编码成字符串
	jsonData, err := json.Marshal(logs)
	if err != nil {
		return err
	}

	// 用AES密钥加密数据
	encryptedData, err := util.AesEncrypt(jsonData, aesKey)
	// encryptedData, err := util.EncryptAES(jsonData, aesKey)
	// util.DecryptAES(encryptedData, aesKey)
	if err != nil {
		return err
	}

	// 将密文发送至服务端API的/logs/uuid
	endpoint := fmt.Sprintf("%s/logs/%s", serverAddress, uuid)
	payload := strings.NewReader(string(encryptedData))
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
