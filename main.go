package main

import (
	"fmt"

	"github.com/FriedCoderZ/LogCollector-client/internal/collect"
)

func main() {
	// config := config.GetConfig()
	// searchPath := config.Collector.SearchPath
	// filePathPattern := config.Collector.FilePathPattern
	// parseTemplate := config.Collector.ParseTemplate
	// serverAddress := config.Server.Address
	// reportInterval := 5
	// database.ClearRecords()
	// collector := collector.NewCollector(searchPath, filePathPattern, parseTemplate, serverAddress, reportInterval)
	// collector.Run()
	err := collect.Register()
	if err != nil {
		fmt.Println(err)
	}
}
