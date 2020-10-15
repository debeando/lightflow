package autoincrement

import (
	"errors"
	"fmt"
	"time"
)

const (
	IntervalbyHours time.Duration = 24
	DateFormat      string        = "2006-01-02"
)

type Duration struct {
	Date time.Time
}

func Date(start string, end string, fn func(string)) error {
	d := Duration{}

	if ok := d.ValidDate(start); !ok {
		return errors.New("Invalid start date.")
	}

	if ok := d.ValidDate(end); !ok {
		return errors.New("Invalid end date.")
	}

	if err := d.Parse(start); err != nil {
		return err
	}

	if d.GreaterThanDate(start, end) {
		return errors.New("Start date should be greater than end date.")
	}

	for {
		fn(d.ToDate())

		if d.ToDate() == end {
			break
		}

		d.Increment(IntervalbyHours)
	}

	return nil
}

func (d *Duration) ValidDate(date string) bool {
	if _, err := time.Parse(DateFormat, date); err != nil {
		return false
	}
	return true
}

func (d *Duration) GreaterThanDate(start string, end string) bool {
	s, _ := time.Parse(DateFormat, start)
	e, _ := time.Parse(DateFormat, end)

	if s.Sub(e).Minutes() > 0 {
		return true
	}

	return false
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
