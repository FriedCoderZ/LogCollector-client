package collector

import (
	"log"
	"path/filepath"
	"time"

	"github.com/FriedCoderZ/LogCollector-client/internal/collect"
	"github.com/FriedCoderZ/LogCollector-client/internal/config"
	"github.com/FriedCoderZ/LogCollector-client/internal/util"
)

type Collector struct {
	SearchPath     string // 待采集文件搜索根目录
	FilePath       string // 待采集文件路径正则表达式
	parseTemplate  string // 解析模板
	ReportInterval int    // 报告间隔
	ServerAddress  string // 服务器地址
	LogPath        string // 日志路径
}

func NewCollector(searchPath, filePath, parseTemplate, serverAddress string, reportInterval int) *Collector {
	searchAbsPath, err := filepath.Abs(searchPath)
	if err != nil {
		log.Fatal(err)
	}
	return &Collector{
		SearchPath:     searchAbsPath,
		FilePath:       filePath,
		parseTemplate:  parseTemplate,
		ReportInterval: reportInterval,
		ServerAddress:  serverAddress,
	}
}

func (c Collector) Run() error {
	config := config.GetConfig()
	for {
		files, err := util.FindAllMatchingFiles(c.SearchPath, c.FilePath)
		if err != nil {
			return err
		}
		var logs []map[string]string
		for _, file := range files {
			logTexts, err := collect.ReadLogByFile(file)
			if err != nil {
				return err
			}
			if len(logTexts) > 0 {
				c.parseTemplate = ReplaceTemplates(c.parseTemplate)
				fileLogs, err := collect.ParseLogs(logTexts, c.parseTemplate)
				if err != nil {
					return err
				}
				logs = append(logs, fileLogs...)
			}
		}
		if len(logs) > 0 {
			log.Printf("新增%d条日志", len(logs))
		}
		if len(logs) > 0 {
			err = collect.SendLogs(logs)
		}
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(config.Collector.ReportInterval) * time.Second)
	}
}
