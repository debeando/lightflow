package variables

import (
	"time"

	"github.com/debeando/lightflow/common"
)

// Items is a collection of map:
type List struct {
	// This variable is the start value to create each combination related
	// by date.
	CurrentTime time.Time
	// Variables to use in the template.
	Items map[string]interface{}
}

var list *List

// Load is a singleton method to return same object:
func Load() *List {
	if list == nil {
		list = &List{
			Items: make(map[string]interface{}),
		}
	}

	return list
}

// Set variables use in the pipe:
func (l *List) Set(variables map[string]interface{}) {
	for index, value := range variables {
		l.Items[index] = value
	}
}

func (l *List) Get(name string) interface{} {
	if value, ok := l.Items[name]; ok {
		return value
	}
	return nil
}

// Verify the variable name exist on the list:
func (l *List) Exist(name string) bool {
	if _, ok := l.Items[name]; ok {
		return true
	}

	return false
}

// Update value on specific variable:
func (l *List) Update(name string, value interface{}) {
	if _, ok := l.Items[name]; ok {
		l.Items[name] = value
	}
}

func (l *List) SetDate(date string) bool {
	new_date := common.StringToDate(date)

	if l.Items["date"] != new_date.Format("2006-01-02") {
		l.Items["date"] = new_date.Format("2006-01-02")
		l.Items["year"] = new_date.Format("2006")
		l.Items["month"] = new_date.Format("01")
		l.Items["day"] = new_date.Format("02")

		l.CurrentTime = new_date

		return true
	}

	return false
}

func (l *List) GetDate() string {
	return common.InterfaceToString(l.Get("date"))
}
