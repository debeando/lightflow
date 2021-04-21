package flow

func (f *Flow) unset() {
	f.Variables.Set(map[string]interface{}{
		"stdout": "",
		"exit_code": 0,
	})

	for _, key := range f.GetPipeUnset() {
		f.Variables.Set(map[string]interface{}{
			key: "",
		})
	}
}
