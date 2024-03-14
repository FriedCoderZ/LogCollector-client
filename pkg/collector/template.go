package collector

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ReplaceTemplates(input string) string {
	// 定义匹配模板的正则表达式
	re := regexp.MustCompile(`\{\{(?P<template>\w+)(?:\((?P<name>\w+)\))?(?::(?P<describe>[^}]+))?\}\}`)

	// 替换匹配到的模板字符串
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		// 获取模板中的 template, name 和 describe
		submatch := re.FindStringSubmatch(match)
		template := submatch[re.SubexpIndex("template")]
		name := submatch[re.SubexpIndex("name")]
		describe := submatch[re.SubexpIndex("describe")]

		// 进行替换操作
		// 替换逻辑可以根据您的需求进行修改
		result, err := replaceByType(template, name, describe)
		if err != nil {
			log.Fatal(err)
		}
		return result
	})

	return output
}

func replaceByType(template, name, describe string) (string, error) {
	switch template {
	case "year":
		return func(template, name, describe string) (string, error) {
			return `(?P<year>d{4})`, nil
		}(template, name, describe)
	case "month":
		return func(template, name, describe string) (string, error) {
			pattern := `(?P<month>(?i)Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)`
			if regexp.MustCompile(pattern).FindStringSubmatch(describe) != nil {
				return pattern, nil
			}
			pattern = `(?P<month>(?i)January|February|March|April|May|June|July|August|September|October|November|December)`
			if regexp.MustCompile(pattern).FindStringSubmatch(describe) != nil {
				return pattern, nil
			}
			return `(?P<month>0?[1-9]|1[0-2])`, nil
		}(template, name, describe)
	case "day":
		return func(template, name, describe string) (string, error) {
			return `(?P<day>0?[1-9]|[1-2][0-9]|3[0-1])`, nil
		}(template, name, describe)
	case "date":
		return func(template, name, describe string) (string, error) {
			if describe == "" {
				return "", fmt.Errorf("please enter a spacer to parse {{date}} e.g. {{date:-}}")
			}
			pattern := []string{`(?P<date>(d{4})`, `(0?[1-9]|1[0-2])`, `(0?[1-9]|[1-2][0-9]|3[0-1]))`}
			return strings.Join(pattern, describe), nil
		}(template, name, describe)
	case "hour":
		return func(template, name, describe string) (string, error) {
			return `(?P<hour>[01]\d|2[0-3])`, nil
		}(template, name, describe)
	case "minute":
		return func(template, name, describe string) (string, error) {
			return `(?P<minute>[0-5]?\d)`, nil
		}(template, name, describe)
	case "second":
		return func(template, name, describe string) (string, error) {
			return `(?P<second>[0-5]?\d)`, nil
		}(template, name, describe)
	case "time":
		return func(template, name, describe string) (string, error) {
			if describe == "" {
				return "", fmt.Errorf("please enter a spacer to parse {{time}} e.g. {{time:-}}")
			}
			pattern := []string{`(?P<time>([01]\d|2[0-3])`, `([0-5]?\d)`, `([0-5]?\d))`}
			return strings.Join(pattern, describe), nil
		}(template, name, describe)
	case "ip", "ipv4":
		return func(template, name, describe string) (string, error) {
			pattern := `(?P<ip>((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))`
			if describe == "mask" {
				pattern = `(?P<ip>((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/(3[0-2]|[1-2]?[0-9]))`
			}
			return pattern, nil
		}(template, name, describe)
	default:
		// 对于未知的 type，直接返回原始模板
		return "", fmt.Errorf("unknown template:" + template)
	}
}
