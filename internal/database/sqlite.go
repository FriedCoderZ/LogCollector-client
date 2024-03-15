package database

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var (
	engine *xorm.Engine
)

func init() {
	// 连接数据库
	fmt.Print("connect to the database...")
	var err error
	engine, err = xorm.NewEngine("sqlite3", "../../data.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("OK\n")
	// 初始化数据库
	fmt.Print("Initialize the database...")
	err = sync()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("OK\n")
}

func sync() error {
	err := engine.Sync2(new(logRecordInfo))
	if err != nil {
		return err
	}
	return nil
}

func GetEngine() *xorm.Engine {
	return engine
}
