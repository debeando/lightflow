package iterator

import (
	"reflect"

	"github.com/debeando/lightflow/flow/duration"
)

type Iterator struct {
	Name          string
	Key           string
	Index         int
	ExecutionTime string
	Items         interface{}
}

func (t *Iterator) Exist(name string) bool {
	for i := range t.Next() {
		if t.key(i) == name {
			return true
		}
	}

	return false
}

func (t *Iterator) Run(fn func()) {
	t.ExecutionTime = duration.Start(func() {
		t.Loops(fn)
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

func (t *Iterator) Loops(fn func()) {
	for t.Index = range t.Next() {
		t.Key = t.key(t.Index)

		if len(t.Name) > 0 {
			if ! t.Exist(t.Name) {
				break
			} else if t.Exist(t.Name) && t.Key != t.Name {
				continue
			}
		}

		fn()
	}
}

func (t *Iterator) key(index int) string {
	items := reflect.ValueOf(t.Items)
	item := items.Index(index)
	if item.Kind() == reflect.Struct {
		return reflect.Indirect(item).Field(0).Interface().(string)
	}

	return ""
}
