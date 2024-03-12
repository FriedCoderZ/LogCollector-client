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

// checkNamedCaptures 函数用于检查正则表达式是否每个捕获组都包含命名
func checkNamedCaptures(pattern string) bool {
	re := regexp.MustCompile(pattern)
	namedCaptures := re.SubexpNames()

	// 遍历命名捕获组，如果有未命名的组则返回false
	for _, name := range namedCaptures {
		if name == "" {
			return false
		}
	}

	return true
}
