package util

import (
	"regexp"
)

// allCapturesNamed 函数用于检查正则表达式是否每个捕获组都包含命名
func allCapturesNamed(pattern string) bool {
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
