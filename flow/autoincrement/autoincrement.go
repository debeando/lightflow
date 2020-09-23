package autoincrement

import (
	"errors"
	"fmt"
	"time"
)

const (
	IntervalbyHours time.Duration = 24
	DateFormat string = "2006-01-02"
)

type Duration struct {
	Date time.Time
}

func Date(start string, end string, fn func(string)) error {
	d := Duration{}

	if ok := d.Valid(start); ! ok {
		return errors.New("Invalid start date.")
	}

	if ok := d.Valid(end); ! ok {
		return errors.New("Invalid end date.")
	}

	if err := d.Parse(start); err != nil {
		return err
	}

	for {
		fn(d.ToDate())

		if (d.ToDate() == end) {
			break
		}

		d.Increment(IntervalbyHours)
	}

	return nil
}

func (d *Duration) Valid(date string) bool {
	if _, err := time.Parse(DateFormat, date); err != nil {
		return false
	}
	return true
}

func (d *Duration) Increment(hour time.Duration) {
	d.Date = d.Date.Add(hour * time.Hour)
}

func (d *Duration) ToDate() string {
	return fmt.Sprint(d.Date.Format(DateFormat))
}

func (d *Duration) Parse(date string) (err error) {
	d.Date, err = time.Parse(DateFormat, date)
	return err
}
