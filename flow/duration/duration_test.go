package duration_test

import (
	"time"
	"testing"

	"github.com/debeando/lightflow/flow/duration"
)

func TestDuration(t *testing.T) {
	et := duration.Start(func() {
		time.Sleep(3 * time.Second)
	})

	if et != "3" {
		t.Errorf("Expected %s, got %s.", "3", et)
	}
}
