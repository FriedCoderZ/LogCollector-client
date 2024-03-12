package collect

import (
	"fmt"

	"github.com/FriedCoderZ/LogCollector-client/internal/util"
)

// ParseLogs 函数根据给定的正则表达式模式解析日志文本，并返回解析后的命名子匹配项列表
func ParseLogs(logTexts []string, pattern string) ([]map[string]string, error) {
	result := []map[string]string{}
	for index, text := range logTexts {
		parsedResult, err := util.ParseString(text, pattern)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log at index %d: %v", index, err)
		}
		result = append(result, parsedResult)
	}
	return result, nil
}
