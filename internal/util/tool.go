package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// ParseString 函数根据正则表达式模式解析文本并返回命名的子匹配项
func ParseString(text string, pattern string) (map[string]string, error) {
	re := regexp.MustCompile(pattern)
	result := make(map[string]string)

	matches := re.FindStringSubmatch(text)
	if matches == nil {
		return result, fmt.Errorf("unable to correctly perform regex grouping")
	}

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result, nil
}

func FindAllMatchingFiles(searchPath string, filePathPattern string) ([]string, error) {
	searchPath, err := filepath.Abs(searchPath)
	if err != nil {
		return nil, err
	}
	absfilePathPattern := "^" + searchPath + "/" + filePathPattern + "$"
	var matchingFiles []string
	// 编译正则表达式
	regExp, err := regexp.Compile(absfilePathPattern)
	if err != nil {
		return nil, err
	}

	// 遍历文件系统
	err = filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		// 忽略根目录
		if path == searchPath {
			return nil
		}
		// 如果路径匹配正则表达式，则添加到结果列表
		if regExp.MatchString(path) {
			matchingFiles = append(matchingFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return matchingFiles, nil
}
