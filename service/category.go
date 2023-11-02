package service

import (
	"errors"
	"strconv"
	"strings"
)

type Category int

const (
	Inbox Category = iota
	_
	Next
	Someday
	Focus
)

var (
	names = [...]string{"Inbox", "", "Next", "Someday", "Focus"}
)

func (c Category) Name() string {
	return names[c]
}

func CategoryFromName(name string) Category {
	switch name {
	case Inbox.Name():
		return Inbox
	case Next.Name():
		return Next
	case Someday.Name():
		return Someday
	case Focus.Name():
		return Focus
	default:
		return -1
	}
}

func (c *Category) UnmarshalJSON(b []byte) error {
	*c = CategoryFromName(strings.Trim(string(b), `"`))
	if *c == -1 {
		return errors.New("invalid category (" + string(b) + ")")
	}
	return nil
}

func (c *Category) MarshalJSON() ([]byte, error) {
	index := int(*c)
	if index < 0 || index >= len(names) {
		return nil, errors.New("invalid category (" + strconv.Itoa(int(*c)) + ")")
	}
	return []byte(c.Name()), nil
}
