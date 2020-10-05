package iterator

import (
	"reflect"

	"github.com/debeando/lightflow/flow/duration"
)

type Iterator struct {
	Index int
	ExecutionTime string
	Items interface{}
}

func (t *Iterator) Exist(name string) bool {
	items := reflect.ValueOf(t.Items)

	if items.Kind() == reflect.Slice {
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)
				if v.Field(0).Interface() == name {
					return true
				}
			}
		}
	}

	return false
}

func (t *Iterator) Run(name string, fn func()) {
	t.ExecutionTime = duration.Start(func(){
		if t.Exist(name) {
			t.One(fn)
		} else {
			t.More(fn)
		}
	})
}

func (t *Iterator) More(fn func()) {
	items := reflect.ValueOf(t.Items)
	if items.Kind() == reflect.Slice {
		for t.Index = 0; t.Index < items.Len(); t.Index++ {
			t.One(fn)
		}
	}
}

func (t *Iterator) One(fn func()) {
	fn()
}
