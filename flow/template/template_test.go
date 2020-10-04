package template_test

import (
	"testing"

	"github.com/swapbyt3s/lightflow/flow/template"
)

func TestRender(t *testing.T) {
	type TestTemplates struct {
		Template string
		Rendered string
		Variables map[string]interface{}
	}

	var testTemplates = map[int]TestTemplates{}
	testTemplates[0] = TestTemplates{Template: "A: {{ .TCA }}"               , Rendered: "A: foo"               , Variables: map[string]interface{}{"TCA": "foo"}}
	testTemplates[1] = TestTemplates{Template: "A: {{ .TCA }}, B: {{ .tcb }}", Rendered: "A: foo, B: bar"       , Variables: map[string]interface{}{"TCA": "foo", "tcb": "bar"}}
	testTemplates[3] = TestTemplates{Template: "A: {{ .TC0 }}, B: {{ .tcb }}", Rendered: "A: <no value>, B: bar", Variables: map[string]interface{}{"tcb": "bar"}}

	for index, _ := range testTemplates {
		rendered, _ := template.Render(testTemplates[index].Template, testTemplates[index].Variables)

		if rendered != testTemplates[index].Rendered {
			t.Errorf("Expected %s, got %s.", testTemplates[index].Rendered, rendered)
		}
	}
}

func TestVariables(t *testing.T) {
}
