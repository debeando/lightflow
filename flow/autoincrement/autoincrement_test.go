package autoincrement_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/autoincrement"
)

func TestAutoincrement(t *testing.T) {
	var counter = 0
	var dates = [...]string{
		"2019-12-30",
		"2019-12-31",
		"2020-01-01",
		"2020-01-02",
		"2020-01-03",
		"2020-01-04",
	}

	autoincrement.Date(
		"2019-12-30",
		"2020-01-04",
		func(date string) {
			if dates[counter] != date {
				t.Errorf("Expected %s, got %s.", dates[counter], date)
			}
			counter++
		})
}

func TestNoAutoincrement(t *testing.T) {
	var counter = 0
	var dates = [...]string{
		"2019-12-30",
	}

	autoincrement.Date(
		"2019-12-30",
		"2019-12-30",
		func(date string) {
			if dates[counter] != date {
				t.Errorf("Expected %s, got %s.", dates[counter], date)
			}
			counter++
		})
}
