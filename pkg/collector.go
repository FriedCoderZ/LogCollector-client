package pkg

import (
	"github.com/FriedCoderZ/LogCollector-client/internal"
)

type collector struct {
	ParseReg       string
	LogPathReg     string
	ReportInterval int
	ServerAddress  string
}

func (c collector) Run() error {
	logs, err := internal.Collect(c.LogPathReg)
	if err != nil {
		return err
	}
	internal.Send(logs)
	return nil
}
