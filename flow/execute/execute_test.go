package execute_test

import (
	"testing"

	"github.com/swapbyt3s/lightflow/flow/execute"
)

func TestExecute(t *testing.T) {
	type TestExecutes struct {
		Command string
		Stdout string
		ExitCode int
	}

	var testExecutes = map[int]TestExecutes{}
	testExecutes[0] = TestExecutes{Command: "echo 'test\n' && exit 0",    Stdout:  "test", ExitCode: 0}
	testExecutes[1] = TestExecutes{Command: "echo -e 'test' && exit 123", Stdout:  "test", ExitCode: 123}
	testExecutes[2] = TestExecutes{Command: "a", Stdout:  "/bin/bash: a: command not found", ExitCode: 127}

	for index, _ := range testExecutes {
		stdout, exit_code := execute.Execute(testExecutes[index].Command)

		if stdout != testExecutes[index].Stdout {
			t.Errorf("Expected %s, got %s.", testExecutes[index].Stdout, stdout)
		}

		if exit_code != testExecutes[index].ExitCode {
			t.Errorf("Expected %d, got %d.", testExecutes[index].ExitCode, exit_code)
		}
	}
}
