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
	for i := range t.Next() {
		if t.Key(i) == name {
			return true
		}
	}

	return false
}

func (t *Iterator) Run(name string, fn func()) {
	t.ExecutionTime = duration.Start(func() {
		t.Loops(name, fn)
	})
}

func (t *Iterator) Next() (<-chan int) {
	chnl := make(chan int)
	go func() {
		items := reflect.ValueOf(t.Items)
		if items.Kind() == reflect.Slice {
			for i := 0; i < items.Len(); i++ {
				chnl <- i
			}
			close(chnl)
		}
	}()
	return chnl
}

func (t *Iterator) Loops(name string, fn func()) {
	for t.Index = range t.Next() {
		t.Name = t.Key(t.Index)

		if len(name) > 0 {
			if ! t.Exist(name) {
				break
			} else if t.Exist(name) && t.Name != name {
				continue
			}
		}

		fn()
	}
}

func (t *Iterator) Key(index int) string {
	items := reflect.ValueOf(t.Items)
	item := items.Index(index)
	if item.Kind() == reflect.Struct {
		return reflect.Indirect(item).Field(0).Interface().(string)
	}

	return ""
}
