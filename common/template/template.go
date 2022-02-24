package template

import (
	// "errors"
	"bytes"
	"regexp"
	"strings"
	"text/template"

	// "github.com/debeando/lightflow/plugins/plugin"
	// "github.com/debeando/lightflow/variables"
)

// type Render struct{
// 	Variable string `yaml:"variable"`
// 	Template string `yaml:"template"`
// }

// func init() {
// 	plugin.Add("Render", func() plugin.Plugin { return &Render{} })
// }

// func (t *Render) Run(event interface{}) (error, bool) {
// 	render, ok := event.(Render)
// 	if !ok {
// 		return errors.New("Invalid struct"), false
// 	}

// 	vars := *variables.Load()

// 	if len(Variables(template.Value)) > 0 {
// 		new_var, err := Render(template.Value, vars.GetItems())
// 		if err != nil {
// 			return err, true
// 		}

// 		vars.Set(map[string]interface{}{
// 			template.Variable: new_var,
// 		})
// 	}

// 	return nil, false
// }

// Render any string with variables.
func Render(textTemplate string, variables map[string]interface{}) (string, error) {
	var b bytes.Buffer

	t, err := template.New("").Option("missingkey=zero").Parse(textTemplate)
	if err != nil {
		return "", err
	}

	if err := t.Execute(&b, ClearEmptyNil(variables)); err != nil {
		return "", err
	}

	return b.String(), nil
}

// Variables is a method to return a list of variables defined into template.
func Variables(textTemplate string) (variables []string) {
	r := regexp.MustCompile(`{{\s*\.([^{}]*)\s*}}`)
	m := r.FindAllStringSubmatch(textTemplate, -1)
	for _, name := range m {
		if len(name) == 2 {
			variables = append(variables, strings.TrimSpace(name[1]))
		}
	}
	return
}

// ClearEmptyNil set nil interface to empty.
func ClearEmptyNil(variables map[string]interface{}) map[string]interface{} {
	for k, v := range variables {
		if v == nil {
			variables[k] = ""
		}
	}

	return variables
}
