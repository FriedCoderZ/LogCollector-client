package database

import (
	"time"
)

type logRecordInfo struct {
	ID           int64  `xorm:"pk autoincr"`
	Path         string `xorm:"unique"`
	LastReadLine int
	LastReadTime time.Time
	CreateAt     time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

type collectorInfo struct {
	UUID      string `xorm:"pk"`
	AESKey    []byte
	CreateAt  time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
