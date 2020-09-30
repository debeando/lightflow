package variables

import (
	"encoding/json"
	"time"

	"github.com/swapbyt3s/lightflow/common"
	"github.com/swapbyt3s/lightflow/common/log"
	"github.com/swapbyt3s/lightflow/config"
)

// Items is a collection of map:
type List struct{
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

	// list.SetDefaults()

	return list
}

// Build standard variables:
func (l *List) SetDefaults() {
	l.Items["path"] = config.Load().General.Temporary_Directory

	l.CurrentTime = time.Now()
	l.Items["date"]  = l.CurrentTime.Format("2006-01-02")
	l.Items["year"]  = l.CurrentTime.Format("2006")
	l.Items["month"] = l.CurrentTime.Format("01")
	l.Items["day"]   = l.CurrentTime.Format("02")
	l.Items["hour"]  = l.CurrentTime.Format("15")

	l.args()

	l.Items["stdout"] = ""
	l.Items["exit_code"] = 0
	l.Items["error"] = ""
	l.Items["status"] = ""
}

// Load variables bypass JSON arguments in the command line:
func (l *List) args() {
	args_vars := common.GetArgVal("variables").(string)

	if len(args_vars) >= 2 {
		err := json.Unmarshal([]byte(args_vars), &l.Items)
		if err != nil {
			log.Warning("Variables", map[string]interface{}{"Message": err})
		}
	}
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
	if new_date := common.StringToDate(date).Format("2006-01-02"); l.Items["date"] != new_date {
		l.Items["date"] = new_date
		return true
	}

	return false
}
