package pkg

import (
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/FriedCoderZ/LogCollector-client/internal/collect"
)

type collector struct {
	ParseReg       string // 解析正则表达式
	FileDirPath    string // 待采集日志文件目录路径
	FileNameReg    string // 待采集日志文件名正则表达式
	ReportInterval int    // 报告间隔
	ServerAddress  string // 服务器地址
	logPath        string // 日志路径
}

func NewCollector(parseReg, fileDirPath, fileNameReg, serverAddress string, reportInterval int) *collector {
	return &collector{
		ParseReg:       parseReg,
		FileDirPath:    fileDirPath,
		FileNameReg:    fileNameReg,
		ReportInterval: reportInterval,
		ServerAddress:  serverAddress,
	}
}

func (c collector) Run() error {
	files, err := ioutil.ReadDir(c.FileDirPath)
	if err != nil {
		return err
	}

	fileNameRegex := regexp.MustCompile(c.FileNameReg)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if fileNameRegex.MatchString(file.Name()) {
			filePath := filepath.Join(c.FileDirPath, file.Name())
			_, err := collect.ReadLog(filePath)
			logTexts, err := collect.ReadLog(filePath)
			logs, err := collect.ParseLogs(logTexts, c.ParseReg)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
