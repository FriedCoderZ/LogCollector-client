package database

import (
	"time"
)

type LogRecord struct {
	ID           int64  `xorm:"pk autoincr"`
	Path         string `xorm:"unique"`
	LastReadLine int
	LastReadTime time.Time
	CreateAt     time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}
