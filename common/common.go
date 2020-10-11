package common

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func GetArgVal(name string) interface{} {
	if flag.Lookup(name) != nil {
		return flag.Lookup(name).Value.(flag.Getter).Get()
	}
	return nil
}

func IsArgStringDefined(name string) bool {
	if len(GetArgVal(name).(string)) == 0 {
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
		for y := l - 1; y >= 0; y-- {
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

func StringToDate(date string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, date)

	return t
}

func StringToInt(value string) int {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	return 0
}

func GetArgValJSON(arg string, key string) (val string, err error) {
	attrs, err := GetArgValsJSON(arg)

	if err != nil {
		return "", err
	}

	return InterfaceToString(attrs[key]), nil
}

func GetArgValsJSON(name string) (attr map[string]interface{}, err error) {
	return StringToJSON(GetArgVal(name).(string))
}

func StringToJSON(text string) (attr map[string]interface{}, err error) {
	if err = json.Unmarshal([]byte(text), &attr); err != nil {
		return map[string]interface{}{}, err
	}

	return attr, nil
}
