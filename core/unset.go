package core

func (core *Core) unset() {
	core.Variables.Set(map[string]interface{}{
		"stdout":    "",
		"exit_code": 0,
	})

	for _, key := range core.GetPipeUnset() {
		core.Variables.Set(map[string]interface{}{
			key: "",
		})
	}
}
