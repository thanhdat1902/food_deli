package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type MyTime time.Time

const MyTimeFormat = "15:04:05"

func (t MyTime) ToString() {
	
}
func NewMyTime(hour, min, sec int) MyTime {

	t := time.Date(0, time.January, 1, hour, min, sec, 0, time.UTC)
	return MyTime(t)
}

func (t *MyTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = MyTime(v)
	case nil:
		*t = MyTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	}
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(MyTimeFormat)), nil
}

func (t *MyTime) UnmarshalText(value string) error {
	dd, err := time.Parse(MyTimeFormat, value)
	if err != nil {
		return err
	}
	*t = MyTime(dd)
	return nil
}

func (MyTime) GormDataType() string {
	return "TIME"
}
