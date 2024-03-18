package main

import (
	"log"

	"github.com/FriedCoderZ/LogCollector-client/internal/collect"
	"github.com/FriedCoderZ/LogCollector-client/internal/config"
	"github.com/FriedCoderZ/LogCollector-client/internal/database"
	"github.com/FriedCoderZ/LogCollector-client/pkg/collector"
)

func main() {
	database.ClearRecords()
	flag, err := database.IsCollectorInfoEmpty()
	if err != nil {
		log.Fatal(err)
	}
	if flag {
		err = collect.Register()
		if err != nil {
			log.Fatal(err)
		}
	}
	config := config.GetConfig()
	searchPath := config.Collector.SearchPath
	filePathPattern := config.Collector.FilePathPattern
	parseTemplate := config.Collector.ParseTemplate
	reportInterval := config.Collector.ReportInterval
	serverAddress := config.Server.Address
	collector := collector.NewCollector(searchPath, filePathPattern, parseTemplate, serverAddress, reportInterval)
	err = collector.Run()
	if err != nil {
		log.Fatal(err)
	}
}
