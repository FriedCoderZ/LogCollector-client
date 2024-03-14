package database

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	// 连接数据库
	fmt.Print("connect to the database...")
	engine, err := xorm.NewEngine("sqlite3", "../../data.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("OK\n")
	Engine = engine
	// 初始化数据库
	fmt.Print("Initialize the database...")
	err = sync()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("OK\n")
}

func sync() error {
	err := Engine.Sync2(new(LogRecord))
	if err != nil {
		return err
	}
	return nil
}
