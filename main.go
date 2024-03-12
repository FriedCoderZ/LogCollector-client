package main

import (
	"fmt"
	"regexp"
)

func parseString(input string, pattern string) map[string]string {
	re := regexp.MustCompile(pattern)
	result := make(map[string]string)

	matches := re.FindStringSubmatch(input)
	if matches == nil {
		return result
	}

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result
}

func main() {
	input := "2024/03/12"
	pattern := `(?P<year>2024/03/12)`

	resultMap := parseString(input, pattern)

	fmt.Println(resultMap)
}
