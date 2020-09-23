package template

// NOTES:
// - Sacar la gestion de errores de aquí.

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"

	"github.com/swapbyt3s/lightflow/common/log"
)

func Render(text_template string, variables map[string]interface{}) string {
	var b bytes.Buffer

	t, err := template.New("").Parse(text_template)
	if err != nil {
		log.Warning("Render", map[string]interface{}{"Message": err.Error()})
	}

	if err := t.Execute(&b, variables); err != nil {
		log.Warning("Render", map[string]interface{}{"Message": err.Error()})
	}

	return b.String()
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
