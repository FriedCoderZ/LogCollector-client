package database

import (
	"fmt"
	"path/filepath"
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

func GetRecord(path string) (*LogRecord, error) {
	var record LogRecord
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	has, err := Engine.Where("path = ?", path).Get(&record)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if has {
		return &record, nil
	} else {
		return nil, nil
	}
}

func CreateRecord(path string) (*LogRecord, error) {
	var record LogRecord
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	record.Path = path
	record.CreateAt = time.Now()
	record.UpdatedAt = time.Now()
	_, err = Engine.Insert(&record)
	if err != nil {
		return &record, err
	}
	return &record, nil
}

func ClearRecords() error {
	_, err := Engine.Exec("DELETE FROM log_record")
	if err != nil {
		return err
	}
	return nil
}

func GetOrCreateRecord(path string) (*LogRecord, error) {
	var record *LogRecord

	record, err := GetRecord(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if record == nil || record.ID == 0 {
		record, err = CreateRecord(path)
		if err != nil {
			return nil, err
		}
	}
	return record, nil
}

func (r *LogRecord) Update(line int, time time.Time) error {
	r.LastReadLine = line
	r.LastReadTime = time
	_, err := Engine.ID(r.ID).Cols("last_read_line", "last_read_time", "updated_at").Update(r)
	return err
}
