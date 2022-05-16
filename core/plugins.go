package core

import (
	// "fmt"
	"reflect"

	"github.com/debeando/lightflow/plugins"
	// "github.com/debeando/lightflow/variables"
)

func (core *Core) Plugins(config interface{}) error {
	items := reflect.ValueOf(config)
	types := items.Type()

	if items.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < items.NumField(); i++ {
		if types.Field(i).Name == "Ignore" {
			if ! IsEmptyPlugin(items.Field(i).Interface()) {
				return nil
			}
		}
	}

	for i := 0; i < items.NumField(); i++ {
		if types.Field(i).Name == "Name" {
			continue
		}
		if types.Field(i).Name == "Evaluate" {
			continue
		}

		if IsEmptyPlugin(items.Field(i).Interface()) {
			continue
		}

		// vars := *variables.Load()
		// vars.Set(map[string]interface{}{
		// 	"pipe": items.Type().Field(i).Name,
		// })

		// Evaluate in every load plugin.
		for y := 0; y < items.NumField(); y++ {
			if items.Type().Field(y).Name == "Evaluate" {
				plugins.Load(
					items.Type().Field(y).Name,
					items.Field(y).Interface(),
				)
			}
		}

		// Load plugin.
		err, repeat := plugins.Load(
			types.Field(i).Name,
			items.Field(i).Interface(),
		)

		if err != nil {
			return err
		}
		if repeat {
			i = 0
		}
	}

	return nil
}

func IsEmptyPlugin(plugin interface{}) bool {
	return reflect.DeepEqual(
		plugin,
		reflect.Zero(reflect.TypeOf(plugin)).Interface(),
	)
}
