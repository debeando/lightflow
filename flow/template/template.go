package template

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

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
		if _, ok := v.(string); !ok {
			variables[k] = ""
		}
	}

	return variables
}
