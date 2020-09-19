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

	// This method are out singleton declaration to build and rebuild
	// variables list.
	list.system()
	list.config()
	list.args()
	// list.register()

	return list
}

// Load variables in the config file:
func (l *List) config() {
	l.Items["path"] = config.Load().General.Temporary_Directory
}

// Build standard variables:
func (l *List) system() {
	l.CurrentTime = time.Now()
	l.Items["date"] = l.CurrentTime.Format("2006-01-02")
	l.Items["year"] = l.CurrentTime.Format("2006")
	l.Items["hour"] = l.CurrentTime.Format("15")
}

// Load variables bypass JSON arguments in the command line:
func (l *List) args() {
	args_vars := common.GetArgVal("variables")

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
