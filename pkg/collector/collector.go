package collector

import (
	"fmt"
	"strings"

	"github.com/FriedCoderZ/LogCollector-client/internal/collect"
	"github.com/FriedCoderZ/LogCollector-client/internal/util"
)

type Collector struct {
	SearchPath      string // 待采集文件搜索根目录
	FilePathPattern string // 待采集文件路径正则表达式(绝对路径)
	parseTemplate   string // 解析模板
	ReportInterval  int    // 报告间隔
	ServerAddress   string // 服务器地址
	LogPath         string // 日志路径
}

func NewCollector(searchPath, filePathPattern, parseTemplate, serverAddress string, reportInterval int) *Collector {
	return &Collector{
		SearchPath:      searchPath,
		FilePathPattern: filePathPattern,
		parseTemplate:   parseTemplate,
		ReportInterval:  reportInterval,
		ServerAddress:   serverAddress,
	}
}

func (c Collector) Run() error {
	// 查找所有匹配的文件
	files, err := util.FindAllMatchingFiles(c.SearchPath, c.FilePathPattern)
	if err != nil {
		return err
	}
	fmt.Printf("找到%d个符合的文件：%s\n", len(files), strings.Join(files, ", "))
	for index, file := range files {
		fmt.Printf("读取第%d个文件：%s\n", index, file)
		logTexts, err := collect.ReadLog(file)
		fmt.Println(logTexts)
		if err != nil {
			return err
		}
		if len(logTexts) > 0 {
			c.parseTemplate = ReplaceTemplates(c.parseTemplate)
			logs, err := collect.ParseLogs(logTexts, c.parseTemplate)
			// Test
			for _, log := range logs {
				fmt.Println(log)
			}
			// End
			if err != nil {
				return err
			}
		}
	}
	return nil
}
