package database

import (
	"errors"
	"path/filepath"
	"time"
)

// LogRecord
func GetRecord(path string) (*logRecordInfo, error) {
	var record logRecordInfo
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	has, err := engine.Where("path = ?", path).Get(&record)
	if err != nil {
		return nil, err
	}
	if has {
		return &record, nil
	} else {
		return nil, nil
	}
}

func CreateRecord(path string) (*logRecordInfo, error) {
	var record logRecordInfo
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	record.Path = path
	record.CreateAt = time.Now()
	record.UpdatedAt = time.Now()
	_, err = engine.Insert(&record)
	if err != nil {
		return &record, err
	}
	return &record, nil
}

func ClearRecords() error {
	_, err := engine.Exec("DELETE FROM log_record_info")
	if err != nil {
		return err
	}
	return nil
}

func GetOrCreateRecord(path string) (*logRecordInfo, error) {
	var record *logRecordInfo

	record, err := GetRecord(path)
	if err != nil {
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

func (r *logRecordInfo) Update(line int, time time.Time) error {
	r.LastReadLine = line
	r.LastReadTime = time
	_, err := engine.ID(r.ID).Cols("last_read_line", "last_read_time", "updated_at").Update(r)
	return err
}

// CollectInfo
func CreateCollectInfo(uuid string, aesKey []byte) (*collectorInfo, error) {
	err := ClearCollectInfo()
	if err != nil {
		return nil, err
	}
	var collectorInfo collectorInfo
	collectorInfo.UUID = uuid
	collectorInfo.AESKey = aesKey
	collectorInfo.CreateAt = time.Now()
	collectorInfo.UpdatedAt = time.Now()
	_, err = engine.Insert(&collectorInfo)
	if err != nil {
		return &collectorInfo, err
	}
	return &collectorInfo, nil
}

func ClearCollectInfo() error {
	_, err := engine.Exec("DELETE FROM collector_info")
	if err != nil {
		return err
	}
	return nil
}

func GetCollectorInfo() (*collectorInfo, error) {
	var collectorInfo collectorInfo
	ok, err := engine.Get(&collectorInfo)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("no any collector info")
	}
	return &collectorInfo, nil
}

func IsCollectorInfoEmpty() (bool, error) {
	count, err := engine.Table("collector_info").Count()
	if err != nil {
		return false, err
	}
	return count == 0, err
}
