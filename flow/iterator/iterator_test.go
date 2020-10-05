package iterator_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/iterator"
)

type Item struct {
	Value string
}

type List struct {
	Name string
	Items []Item
}

var demo = []List{
	{ Name: "foo"},
	{ Name: "bar"},
	{ Name: "baz"},
}

func TestExistOk(t *testing.T) {
	itr := iterator.Iterator{
		Items: demo,
	}

	t.Log(itr.Exist("ba"))
}

func TestExistKo(t *testing.T) {
	itr := iterator.Iterator{
		Items: demo,
	}

	t.Log(itr.Exist("bar"))
}

func TestLoop(t *testing.T) {
	itr := iterator.Iterator{
		Items: demo,
	}

	itr.Run("", func() {
		t.Log(itr.Index)
	})
}
