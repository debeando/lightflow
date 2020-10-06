package iterator_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/iterator"
)

type List struct {
	Name  string
}

var demo = []List{
	{Name: "foo"},
	{Name: "bar"},
	{Name: "baz"},
}

func TestExistOk(t *testing.T) {
	itr := iterator.Iterator{
		Items: demo,
	}

	if exist := itr.Exist("ba"); exist != false {
		t.Errorf("Expected %t, got %t.", false, exist)
	}
}

func TestExistKo(t *testing.T) {
	itr := iterator.Iterator{
		Items: demo,
	}

	if exist := itr.Exist("baz"); exist != true {
		t.Errorf("Expected %t, got %t.", true, exist)
	}
}

func TestLoop(t *testing.T) {
	counter := 0
	itr := iterator.Iterator{
		Items: demo,
	}

	itr.Run("", func() {
		counter++
	})

	if counter != len(demo) {
		t.Errorf("Expected %d, got %d.", len(demo), counter)
	}
}
