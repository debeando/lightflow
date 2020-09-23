package template

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

func Render(text_template string, variables map[string]interface{}) (string, error) {
	var b bytes.Buffer

	t, err := template.New("").Parse(text_template)
	if err != nil {
		return "", err
	}

	if err := t.Execute(&b, variables); err != nil {
		return "", err
	}

	return b.String(), nil
}

func Variables(text_template string) (variables []string) {
	r := regexp.MustCompile(`{{\s*\.([^{}]*)\s*}}`)
	m := r.FindAllStringSubmatch(text_template, -1)
	for _, name := range m {
		if len(name) == 2 {
			variables = append(variables, strings.TrimSpace(name[1]))
		}
	}
	return
}
