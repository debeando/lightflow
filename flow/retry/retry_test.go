package retry_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/retry"
)

func TestRetryOK(t *testing.T) {
	var counter = 0

	r := retry.Retry{
		Attempt: 3,
		Wait:    1,
	}

	r.Retry(
		func() bool {
			counter++
			return true
		})

	if counter != 3 {
		t.Errorf("Expected %d, got %d.", 3, counter)
	}
}

func TestRetryKO(t *testing.T) {
	var counter = 0

	r := retry.Retry{
		Attempt: 3,
		Wait:    1,
	}

	r.Retry(
		func() bool {
			counter++
			return false
		})

	if counter != 1 {
		t.Errorf("Expected %d, got %d.", 1, counter)
	}
}
