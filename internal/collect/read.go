package collect

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLog 从指定的日志文件中按行读取日志数据，并返回一个字符串列表。
func ReadLog(filePath string) ([]string, error) {
	// 打开日志文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 创建字符串列表，用于存储日志数据
	logs := []string{}

	// 创建一个带缓冲的读取器
	scanner := bufio.NewScanner(file)

	// 逐行读取日志数据
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	// 检查读取过程中是否发生错误
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}

	// 返回日志数据字符串列表
	return logs, nil
}
