package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// HTime 自定义时间类型
type HTime struct {
	time.Time
}

// 使用一个具体的时间来定义格式
var (
	formatTime = "2006-01-02 15:04:05"
)

// MarshalJSON 通过json.Marshal来序列化包含HTime的结构体时，会使用自定义的时间输出
func (t HTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

// UnmarshalJSON 通过json.Unmarshal来反序列化包含HTime的结构体时，会使用自定义的时间输入
func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	*t = HTime{Time: now}
	return
}

// Value 写入数据库前进行转换
func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 从数据库读取时进行转换
func (t *HTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
