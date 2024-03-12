package util

import (
	"fmt"
	"regexp"
)

// ParseString 函数根据正则表达式模式解析文本并返回命名的子匹配项
func ParseString(text string, pattern string) (map[string]string, error) {
	re := regexp.MustCompile(pattern)
	result := make(map[string]string)

	matches := re.FindStringSubmatch(text)
	if matches == nil {
		return result, nil
	}

	for i, name := range re.SubexpNames() {
		if name == "" {
			return map[string]string{}, fmt.Errorf("failed to parse named subexpression at index %d", i)
		}
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result, nil
}
