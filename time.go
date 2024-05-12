package objects

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const (
	DATE_FORMAT_DATE            = "2006-01-02"
	DATE_FORMAT_DATETIME        = "2006-01-02 15:04:05"
	DATE_FORMAT_DATEMINUTE      = "2006-01-02 15:04"
	DATE_FORMAT_DATE_CN         = "2006年01月02日"
	DATE_FORMAT_MINIFY_DATE     = "20060102"
	DATE_FORMAT_MINIFY_MONTH    = "200601"
	DATE_MINIFY_FORMAT_DATETIME = "20060102150405"
)

type (
	Date           time.Time
	TimeDateTime   time.Time
	TimeDateMinute time.Time
	TimeDateCN     time.Time

	TimeFormat struct {
		Layout string
	}

	TimeX struct {
		Time   time.Time
		Format *TimeFormat
	}
)

func (this Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(DATE_FORMAT_DATE))
	return []byte(stamp), nil
}

func (td *Date) UnmarshalJSON(b []byte) error {
	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return err
	}
	t, err := time.Parse(DATE_FORMAT_DATE, timeStr)
	if err != nil {
		return err
	}
	*td = Date(t)
	return nil
}
func (t Date) Value() (driver.Value, error) {
	var zeroTime time.Time
	_t := time.Time(t)
	if _t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return _t.Format(DATE_FORMAT_DATE), nil
}

func (t *Date) Scan(value interface{}) error {
	if value == nil {
		*t = Date(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = Date(v)
		return nil
	case []byte:
		tt, err := time.Parse(DATE_FORMAT_DATE, string(v))
		if err != nil {
			return err
		}
		*t = Date(tt)
		return nil
	case string:
		tt, err := time.Parse(DATE_FORMAT_DATE, v)
		if err != nil {
			return err
		}
		*t = Date(tt)
		return nil
	default:
		return fmt.Errorf("can't convert %T to TimeDateTime", value)
	}
}

func (this TimeDateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(DATE_FORMAT_DATETIME))
	return []byte(stamp), nil
}
func (td *TimeDateTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return err
	}
	t, err := time.Parse(DATE_FORMAT_DATETIME, timeStr)
	if err != nil {
		return err
	}
	*td = TimeDateTime(t)
	return nil
}
func (t TimeDateTime) Value() (driver.Value, error) {
	// 将 TimeDateTime 类型的值转换为 time.Time 类型的值
	var zeroTime time.Time
	_t := time.Time(t)
	if _t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return _t, nil
}

func (t *TimeDateTime) Scan(value interface{}) error {
	if value == nil {
		*t = TimeDateTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = TimeDateTime(v)
		return nil
	case []byte:
		tt, err := time.Parse(DATE_FORMAT_DATETIME, string(v))
		if err != nil {
			return err
		}
		*t = TimeDateTime(tt)
		return nil
	case string:
		tt, err := time.Parse(DATE_FORMAT_DATETIME, v)
		if err != nil {
			return err
		}
		*t = TimeDateTime(tt)
		return nil
	default:
		return fmt.Errorf("can't convert %T to TimeDateTime", value)
	}
}

func (this TimeDateMinute) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(DATE_FORMAT_DATEMINUTE))
	return []byte(stamp), nil
}

func (this TimeDateCN) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format(DATE_FORMAT_DATE_CN))
	return []byte(stamp), nil
}

func (tf *TimeFormat) MarshalTime(t time.Time) ([]byte, error) {
	return json.Marshal(t.Format(tf.Layout))
}

func (tf *TimeFormat) UnmarshalTime(b []byte) (time.Time, error) {
	var timeStr string
	if err := json.Unmarshal(b, &timeStr); err != nil {
		return time.Time{}, err
	}
	return time.Parse(tf.Layout, timeStr)
}

func (ct *TimeX) MarshalJSON() ([]byte, error) {
	return ct.Format.MarshalTime(ct.Time)
}

func (ct *TimeX) UnmarshalJSON(b []byte) error {
	t, err := ct.Format.UnmarshalTime(b)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}
