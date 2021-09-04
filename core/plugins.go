package core

import (
	"reflect"

	"github.com/debeando/lightflow/plugins"
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

		if IsEmptyPlugin(items.Field(i).Interface()) {
			continue
		}

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
