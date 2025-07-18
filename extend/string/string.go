package cherryString

import (
	"strconv"
	goStrings "strings"

	jsoniter "github.com/json-iterator/go"
)

// CutLastString 截取字符串中最后一段，以@beginChar开始,@endChar结束的字符
// @text 文本
// @beginChar 开始
func CutLastString(text, beginChar, endChar string) string {
	if text == "" || beginChar == "" || endChar == "" {
		return ""
	}

	textRune := []rune(text)

	beginIndex := goStrings.LastIndex(text, beginChar)
	endIndex := goStrings.LastIndex(text, endChar)
	if endIndex < 0 || endIndex < beginIndex {
		endIndex = len(textRune)
	}

	return string(textRune[beginIndex+1 : endIndex])
}

func IsBlank(value string) bool {
	return value == ""
}

func IsNotBlank(value string) bool {
	return value != ""
}

func ToBool(value string) (bool, bool) {
	val, err := strconv.ParseBool(value)
	if err != nil {
		return false, false
	}
	return val, true
}

func ToBoolD(value string) bool {
	val, _ := strconv.ParseBool(value)
	return val
}

func ToUint(value string, def ...uint) (uint, bool) {
	val, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return uint(val), true
}

func ToUintD(value string, def ...uint) uint {
	val, _ := ToUint(value, def...)
	return val
}

func ToInt(value string, def ...int) (int, bool) {
	val, err := strconv.Atoi(value)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return val, true
}

func ToIntD(value string, def ...int) int {
	val, _ := ToInt(value, def...)
	return val
}

func ToInt32(value string, def ...int32) (int32, bool) {
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return int32(val), true
}

func ToInt32D(value string, def ...int32) int32 {
	val, _ := ToInt32(value, def...)
	return val
}

func ToInt64(value string, def ...int64) (int64, bool) {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return val, true
}

func ToInt64D(value string, def ...int64) int64 {
	val, _ := ToInt64(value, def...)
	return val
}

func ToFloat64(value string, def ...float64) (float64, bool) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return f, true
}

func ToFloat64D(value string, def ...float64) float64 {
	f, _ := ToFloat64(value, def...)
	return f
}

func ToFloat32(value string, def ...float32) (float32, bool) {
	f64, ok := ToFloat64(value)
	if !ok {
		if len(def) > 0 {
			return def[0], false
		}
		return 0, false
	}
	return float32(f64), true
}

func ToFloat32D(value string, def ...float32) float32 {
	f, _ := ToFloat32(value, def...)
	return f
}

func ToString(value interface{}) string {
	ret := ""

	if value == nil {
		return ret
	}

	switch t := value.(type) {
	case string:
		ret = t
	case bool:
		ret = strconv.FormatBool(t)
	case int:
		ret = strconv.Itoa(t)
	case int32:
		ret = strconv.Itoa(int(t))
	case int64:
		ret = strconv.FormatInt(t, 10)
	case uint:
		ret = strconv.Itoa(int(t))
	case uint32:
		ret = strconv.Itoa(int(t))
	case uint64:
		ret = strconv.Itoa(int(t))
	default:
		v, _ := jsoniter.Marshal(t)
		ret = string(v)
	}

	return ret
}

func ToStringSlice(val []interface{}) []string {
	var result []string
	for _, item := range val {
		v, ok := item.(string)
		if ok {
			result = append(result, v)
		}
	}
	return result
}

func SplitIndex(s, sep string, index int) (string, bool) {
	ret := goStrings.Split(s, sep)
	if index >= len(ret) {
		return "", false
	}
	return ret[index], true
}
