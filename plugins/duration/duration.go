package duration

import (
	"fmt"
	"time"
)

func Start(fn func()) string {
	t1 := time.Now()
	fn()
	t2 := time.Now()
	diff := t2.Sub(t1)
	out := time.Time{}.Add(diff)

	return fmt.Sprint(out.Format("15:04:05"))
}
