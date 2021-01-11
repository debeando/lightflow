package date_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/date"
)

func TestValidDate(t *testing.T) {
	type TestDate struct {
		Date string
		Valid bool
	}

	var ai = date.Duration{}
	var testDates = map[int]TestDate{}

	testDates[0] = TestDate{Date: "2019-12-31", Valid: true}
	testDates[1] = TestDate{Date: "2019-02-31", Valid: false}
	testDates[2] = TestDate{Date: "2019-2-27",  Valid: false}

	for index, _ := range testDates {

		if v := ai.ValidDate(testDates[index].Date); v != testDates[index].Valid {
			t.Errorf("Expected %t, got %t.", testDates[index].Valid, v)
		}
	}
}

func TestIncrement(t *testing.T) {
	var counter = 0
	var dates = [...]string{
		"2019-12-30",
		"2019-12-31",
		"2020-01-01",
		"2020-01-02",
		"2020-01-03",
		"2020-01-04",
	}

	date.Increment(
		"2019-12-30",
		"2020-01-04",
		func(date string) {
			if dates[counter] != date {
				t.Errorf("Expected %s, got %s.", dates[counter], date)
			}
			counter++
		})
}

func TestDecrement(t *testing.T) {
	var counter = 0
	var dates = [...]string{
		"2020-01-02",
		"2020-01-01",
		"2019-12-31",
		"2019-12-30",
		"2019-12-29",
	}

	date.Decrement(
		"2020-01-02",
		"2019-12-29",
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

	date.Increment(
		"2019-12-30",
		"2019-12-30",
		func(date string) {
			if dates[counter] != date {
				t.Errorf("Expected %s, got %s.", dates[counter], date)
			}
			counter++
		})
}

func TestGreaterThanDate(t *testing.T) {
	type TestGreaterThanDate struct {
		Start string
		End string
		Valid bool
	}

	var ai = date.Duration{}
	var testGreaterThanDates = map[int]TestGreaterThanDate{}

	testGreaterThanDates[0] = TestGreaterThanDate{Start: "2019-12-31", End: "2019-12-31", Valid: false}
	testGreaterThanDates[1] = TestGreaterThanDate{Start: "2019-11-30", End: "2019-12-01", Valid: false}
	testGreaterThanDates[2] = TestGreaterThanDate{Start: "2019-10-10", End: "2018-10-10", Valid: true}

	for index, _ := range testGreaterThanDates {
		if v := ai.GreaterThanDate(testGreaterThanDates[index].Start, testGreaterThanDates[index].End); v != testGreaterThanDates[index].Valid {
			t.Errorf("Expected %t, got %t.", testGreaterThanDates[index].Valid, v)
		}
	}
}

func TestLessThanDate(t *testing.T) {
	type TestLessThanDate struct {
		Start string
		End string
		Valid bool
	}

	var ai = date.Duration{}
	var testLessThanDates = map[int]TestLessThanDate{}

	testLessThanDates[0] = TestLessThanDate{Start: "2019-12-31", End: "2019-12-31", Valid: false}
	testLessThanDates[1] = TestLessThanDate{Start: "2019-11-30", End: "2019-12-01", Valid: true}
	testLessThanDates[2] = TestLessThanDate{Start: "2019-10-10", End: "2018-10-10", Valid: false}

	for index, _ := range testLessThanDates {
		if v := ai.LessThanDate(testLessThanDates[index].Start, testLessThanDates[index].End); v != testLessThanDates[index].Valid {
			t.Errorf("Expected %t, got %t.", testLessThanDates[index].Valid, v)
		}
	}
}
