package main

import (
	"github.com/FriedCoderZ/LogCollector-client/internal/config"
	"github.com/FriedCoderZ/LogCollector-client/internal/database"
	"github.com/FriedCoderZ/LogCollector-client/pkg/collector"
)

func main() {
	config := config.GetConfig()
	searchPath := config.Collector.SearchPath
	filePathPattern := config.Collector.FilePathPattern
	parseTemplate := config.Collector.ParseTemplate
	serverAddress := config.Server.Address
	// searchPath := "/ouryun/LogCollector-client/log"
	// filePathPattern := `^\/ouryun\/LogCollector-client\/log\/(\d{8})\/service.log$`
	// parseTemplate := `^{{month:Mar}} {{day}} {{time::}} (?P<Host>\w+) (?P<Text>.*)$`
	// serverAddress := "http://example.com"
	reportInterval := 5
	database.ClearRecords()
	collector := collector.NewCollector(searchPath, filePathPattern, parseTemplate, serverAddress, reportInterval)
	collector.Run()
	// parseTemplate := `^{{month:Mar}} (?P<Text>.*)$`
	// res := collector.ReplaceTemplates(parseTemplate)
	// fmt.Println(res)
}
