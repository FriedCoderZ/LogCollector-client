package collect

import (
	"log"
	"strconv"
	"time"

	"github.com/FriedCoderZ/LogCollector-client/internal/util"
)

// ParseLogs 函数根据给定的正则表达式模式解析日志文本，并返回解析后的命名子匹配项列表
func ParseLogs(logTexts []string, pattern string) ([]map[string]string, error) {
	result := []map[string]string{}
	for index, text := range logTexts {
		parsedResult, err := util.ParseString(text, pattern)
		if err != nil {
			log.Printf("failed to parse log at index %d: %v\n无法解析的行内容:%s", index, err, text)
			continue
		}
		result = append(result, parsedResult)
	}
	calcTimestamp(result)
	return result, nil
}

func calcTimestamp(data []map[string]string) {
	currentTime := time.Now()
	for _, item := range data {
		year, _ := strconv.Atoi(item["year"])
		if year == 0 {
			year = currentTime.Year()
		}

		month, _ := strconv.Atoi(item["month"])
		if month == 0 {
			month = int(currentTime.Month())
		}

		day, _ := strconv.Atoi(item["day"])
		if day == 0 {
			day = currentTime.Day()
		}

		hour, _ := strconv.Atoi(item["hour"])
		if hour == 0 {
			hour = currentTime.Hour()
		}

		minute, _ := strconv.Atoi(item["minute"])
		if minute == 0 {
			minute = currentTime.Minute()
		}

		second, _ := strconv.Atoi(item["second"])
		if second == 0 {
			second = currentTime.Second()
		}

		timestamp := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local).Unix()
		item["timestamp"] = strconv.FormatInt(timestamp, 10)
	}
}
