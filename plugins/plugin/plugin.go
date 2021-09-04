package plugin

// Input defines the interface that can interact with the registry
type Plugin interface {
	Run(event interface{}) (error, bool)
}

// Creator lets us use a closure to get intsances of the Input struct
type Creator func() Plugin

// Plugins registry
var Plugins = map[string]Creator{}

// Add can be called from init() on a plugin in this package
// It will automatically be added to the Plugins map to be called externally
func Add(name string, creator Creator) {
	Plugins[name] = creator
}
