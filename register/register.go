package register

import (
	"sync"
)

// Variables is a collection of map:
type Register struct{
	mu     sync.Mutex
	Values map[string]string
}

// var register *Register:
var list *Register

// Load is a singleton method to return same object:
func Load() *Register {
	if list == nil {
		list = &Register{
			Values: make(map[string]string),
		}
	}
	return list
}

// Save or update item on the regist:
func (r *Register) Save(variable string, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Values[variable] = value
}

// Get all items on the regist:
func (r *Register) Get() map[string]string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.Values
}
