package service

import (
	"strings"
	"time"
)

type TaskTime time.Time

func (t *TaskTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	tt, err := time.Parse("20060102T150405Z", value)
	if err != nil {
		return err
	}
	*t = TaskTime(tt)
	return nil
}

func (t TaskTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format("20060102T150405Z")), nil
}
