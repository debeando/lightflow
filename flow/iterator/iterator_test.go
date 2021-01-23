package iterator_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/iterator"
)

type List struct {
	Name   string
	Ignore bool
	Items  []Item
}

type Item struct {
	Name  string
	Value string
}

var demo = []List{
	{
		Name:   "foo",
		Ignore: true,
		Items: []Item{
			{Name: "L1I1"},
			{Name: "L1I2"},
			{Name: "L1I3"},
			{Name: "L1I4"},
		},
	},
	{
		Name: "bar",
		Items: []Item{
			{Name: "L2I1"},
			{Name: "L2I2"},
			{Name: "L2I3"},
		},
	},
	{
		Name: "baz",
		Items: []Item{
			{Name: "L3I1"},
			{Name: "L3I2"},
		},
	},
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

func TestLoopCount(t *testing.T) {
	counter := 0
	itr := iterator.Iterator{
		Items: demo,
	}

	itr.Run(func() {
		counter++
	})

	if counter != 2 {
		t.Errorf("Expected %d, got %d.", 2, counter)
	}
}

func TestLevelOne(t *testing.T) {
	counter := 0
	itrl1 := iterator.Iterator{
		Items: demo,
	}

	itrl1.Run(func() {
		itrl2 := iterator.Iterator{
			Items: demo[itrl1.Index].Items,
		}

		itrl2.Run(func() {
			counter++
		})
	})

	if counter != 5 {
		t.Errorf("Expected %d, got %d.", 5, counter)
	}
}

func TestLevelOneMatchItemName(t *testing.T) {
	counter := 0
	itrl1 := iterator.Iterator{
		Items: demo,
		Name:  "bar",
	}

	itrl1.Run(func() {
		itrl2 := iterator.Iterator{
			Items: demo[itrl1.Index].Items,
			Name:  "L2I2",
		}

		itrl2.Run(func() {
			counter++

			if itrl2.Index != 1 {
				t.Errorf("Expected %d, got %d.", 1, itrl2.Index)
			}
		})
	})

	if counter != 1 {
		t.Errorf("Expected %d, got %d.", 1, counter)
	}
}

func TestLevelOneMatchLoopName(t *testing.T) {
	counter := 0
	itrl1 := iterator.Iterator{
		Items: demo,
		Name:  "bar",
	}

	itrl1.Run(func() {
		itrl2 := iterator.Iterator{
			Items: demo[itrl1.Index].Items,
		}

		itrl2.Run(func() {
			counter++
		})
	})

	if counter != len(demo[1].Items) {
		t.Errorf("Expected %d, got %d.", len(demo[1].Items), counter)
	}
}

func TestIgnore(t *testing.T) {
	counter := 0
	itrl1 := iterator.Iterator{
		Items: demo,
		Name:  "foo",
	}

	itrl1.Run(func() {
		counter++
	})

	if counter != 0 {
		t.Errorf("Expected %d, got %d.", 0, counter)
	}
}
