package iterator

import (
	"reflect"

	"github.com/debeando/lightflow/flow/duration"
)

type Iterator struct {
	Name          string
	Index         int
	ExecutionTime string
	Items         interface{}
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
	t.ExecutionTime = duration.Start(func() {
		t.Loops(name, fn)
	})
}

func (t *Iterator) Loops(name string, fn func()) {
	items := reflect.ValueOf(t.Items)
	if items.Kind() == reflect.Slice {
		for t.Index = 0; t.Index < items.Len(); t.Index++ {
			t.Name = t.Key()

			if len(name) > 0 && ! t.Exist(name) {
				break
			} else if len(name) > 0 && t.Exist(name) && t.Name != name {
				continue
			}

			fn()
		}
	}
}

func (t *Iterator) Key() string {
	items := reflect.ValueOf(t.Items)
	item := items.Index(t.Index)
	if item.Kind() == reflect.Struct {
		return reflect.Indirect(item).Field(0).Interface().(string)
	}

	return ""
}
