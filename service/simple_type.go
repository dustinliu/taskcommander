package service

import (
	"fmt"
	"slices"
	"sync"
)

type Status uint8

const (
	StatusTodo Status = iota
	StatusDone
)

//var StatusMap = map[Status]string{
//StatusTodo: "Todo",
//StatusDone: "Done",
//}

var statusList = []simpleType[Status]{
	{StatusTodo, "Todo"},
	{StatusDone, "Done"},
}

var (
	statusOnce sync.Once
	statuses   []Status
)

func Statuses() []Status {
	statusOnce.Do(func() {
		statuses = getSimpleTypes(statusList)
	})
	return statuses
}

func (s Status) IsValid() bool {
	return isValidSimpleType(s, statusList)
}

func (s Status) Name() string {
	return getSimpleTypeName(s, statusList)
}

type Category uint8

const (
	CategoryInbox Category = iota
	CategoryNext
	CategorySomeday
	CategoryFocus
)

//var categorieyMap = map[Category]string{
//CategoryInbox:   "Inbox",
//CategoryNext:    "Next",
//CategorySomeday: "Someday",
//CategoryFocus:   "Focus",
//}

var categoryList = []simpleType[Category]{
	{CategoryInbox, "Inbox"},
	{CategoryNext, "Next"},
	{CategorySomeday, "Someday"},
	{CategoryFocus, "Focus"},
}

var (
	categories = []Category{}
	catOnce    sync.Once
)

func Categories() []Category {
	catOnce.Do(func() {
		categories = getSimpleTypes(categoryList)
	})
	return categories
}

type simpleType[T ~uint8] struct {
	Value T
	Name  string
}

func (c Category) IsValid() bool {
	return isValidSimpleType(c, categoryList)
}

func (c Category) Name() string {
	return getSimpleTypeName(c, categoryList)
}

func getSimpleTypes[T ~uint8](l []simpleType[T]) []T {
	keys := make([]T, 0, len(l))
	for _, s := range l {
		keys = append(keys, s.Value)
	}
	return keys
}

func isValidSimpleType[T ~uint8](t T, l []simpleType[T]) bool {
	return slices.ContainsFunc(l, func(s simpleType[T]) bool { return s.Value == t })
}

func getSimpleTypeName[T ~uint8](t T, l []simpleType[T]) string {
	if !isValidSimpleType(t, l) {
		return fmt.Sprintf("Invalid %T(%d)", t, t)
	}

	return l[slices.IndexFunc(l, func(s simpleType[T]) bool { return s.Value == t })].Name
}
