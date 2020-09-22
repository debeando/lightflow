package common

import (
	"time"
	"fmt"
	"flag"
)

func GetArgVal(name string) string {
	if flag.Lookup(name) != nil {
		return flag.Lookup(name).Value.(flag.Getter).Get().(string)
	}
	return ""
}

func IsArgDefined(name string) bool {
	if GetArgVal(name) == "" {
		return false
	}

	return true
}

func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func TrimNewlines(text string) string {
	t := []rune(text)
	l := len(t)
	r := []rune("")

	Next := func() int {
		for y := 0; y < l; y++ {
			if t[y] != '\n' {
				return y
			}
		}
		return 0
	}

	Previous := func() int {
		for y := l-1; y >= 0; y-- {
			if t[y] != '\n' {
				return y
			}
		}

		return 0
	}

	for x := 0; x < l; x++ {
		if Next() > x || Previous() < x {
			continue
		}

		r = append(r, t[x])
	}

	return string(r)
}

func Duration(fn func()) string {
	t1 := time.Now()
	fn()
	t2 := time.Now()
	diff := t2.Sub(t1)
	out := time.Time{}.Add(diff)

	return fmt.Sprintln(out.Format("15:04:05"))
}
