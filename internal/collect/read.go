package collect

import (
	"bufio"
	"os"
	"path/filepath"
	"time"

	"github.com/FriedCoderZ/LogCollector-client/internal/database"
)

// ReadLogByFile 从指定的日志文件中按行读取日志数据，并返回一个字符串列表。
func ReadLogByFile(path string) ([]string, error) {
	//更新为全局变量
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	record, err := database.GetOrCreateRecord(path)
	if err != nil {
		return nil, err
	}
	// 获取日志文件的信息
	fileInfo, err := os.Stat(record.Path)
	if err != nil {
		return nil, err
	}
	// 判断文件最后修改时间是否超过上次读取时间
	if !fileInfo.ModTime().After(record.LastReadTime) {
		return nil, nil // 文件未修改，无需读取新增内容
	}

	startLine := record.LastReadLine + 1
	lines, err := readLogByLine(record.Path, startLine)
	if err != nil {
		return nil, err
	}

	// 更新LogRecord对象的LastReadLine和LastReadTime
	record.Update(record.LastReadLine+len(lines), time.Now())
	return lines, nil
}

func readLogByLine(filename string, startLine int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 定位到指定行
	for i := 1; i < startLine; i++ {
		if !scanner.Scan() {
			// 如果文件行数不足，则返回空切片
			return []string{}, nil
		}
	}

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
