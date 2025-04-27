// Package cherrySlice code from: https://github.com/beego/beego/blob/develop/core/utils/slice.go
package cherrySlice

import (
	"math/rand"
	"reflect"
	"strings"

	cstring "github.com/cherry-game/cherry/extend/string"
	cutils "github.com/cherry-game/cherry/extend/utils"
)

// Addable 定义一个接口，表示支持 + 操作的类型
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// Int32In 判定int是否在切片中
func Int32In(v int32, sl []int32) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

// Int64In 判定int64是否在切片中
func Int64In(v int64, sl []int64) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

// StringIn 判定字符串是否在切片中
func StringIn(v string, sl []string) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

// In 判定值是否在切片中
func In[V comparable](v V, sl []V) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

// RandList 生成一个包含[minValue, maxValue]中所有数字的随机数切片
func RandList(minValue, maxValue int) []int {
	if maxValue < minValue {
		minValue, maxValue = maxValue, minValue
	}

	length := maxValue - minValue + 1
	list := rand.Perm(length)
	for index := range list {
		list[index] += minValue
	}
	return list
}

// Merge 合并两个切片
func Merge[V any](slice1, slice2 []V) (c []V) {
	c = append(slice1, slice2...)
	return
}

// Reduce 通过reduce函数解析每个值后生成一个新片
func Reduce[V any](slice []V, a func(V) V) (destSlice []V) {
	for _, v := range slice {
		destSlice = append(destSlice, a(v))
	}
	return
}

// Rand 从切片中随机返回一个值(如果切片为空，返回V类型的默认值)
func Rand[V any](a []V) (b V) {
	length := len(a)
	if length == 0 {
		return
	}

	randNum := rand.Intn(len(a))
	b = a[randNum]
	return
}

// Sum 对切片所有元素求和
func Sum[V Addable](intSlice []V) (sum V) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// Filter 获取切片中满足条件的值
func Filter[V any](slice []V, a func(V) bool) (filterSlice []V) {
	for _, v := range slice {
		if a(v) {
			filterSlice = append(filterSlice, v)
		}
	}
	return
}

// Diff 求切片1中不在切片2中的值
func Diff[V comparable](slice1, slice2 []V) (diffSlice []V) {
	for _, v := range slice1 {
		if _, found := In(v, slice2); !found {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// Intersect 求切片1和切片2的交集
func Intersect[V comparable](slice1, slice2 []V) (sameSlice []V) {
	for _, v := range slice1 {
		if _, found := In(v, slice2); found {
			sameSlice = append(sameSlice, v)
		}
	}
	return
}

// Chunk 将切片分成指定大小的多个切片
func Chunk[V any](slice []V, size int) (chunkSlice [][]V) {
	if size >= len(slice) {
		chunkSlice = append(chunkSlice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkSlice = append(chunkSlice, slice[i:end])
		end += size
	}
	return
}

// Range 生成一个从start到end范围指定步长下标切片
func Range(start, end, step int) (indexSlice []int) {
	for i := start; i <= end; i += step {
		indexSlice = append(indexSlice, i)
	}
	return
}

// Pad 扩展切片到指定长度，用指定值来扩展
func Pad[V any](slice []V, size int, val V) []V {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

// Uniques 将多个切片去重
func Uniques[T comparable](slices ...[]T) []T {
	keys := map[T]struct{}{}

	for _, slice := range slices {
		for _, s := range slice {
			keys[s] = struct{}{}
		}
	}

	var uniqueSlice []T

	for t := range keys {
		uniqueSlice = append(uniqueSlice, t)
	}

	return uniqueSlice
}

// Unique 将一个切片去重
func Unique[T comparable](slice ...T) []T {
	return Uniques[T](slice)
}

// Shuffle 将切片打乱
func Shuffle[V any](slice []V) []V {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

// StringToInt 将字符串切片转换为int切片(不可转的跳过)
func StringToInt(strSlice []string) []int {
	var intSlice []int

	for _, s := range strSlice {
		if cutils.IsNumeric(s) {
			val, ok := cstring.ToInt(s)
			if ok {
				intSlice = append(intSlice, val)
			}
		}
	}

	return intSlice
}

// StringToInt32 将字符串切片转换为int32切片(不可转的跳过)
func StringToInt32(strSlice []string) []int32 {
	var intSlice []int32

	for _, s := range strSlice {
		if cutils.IsNumeric(s) {
			val, ok := cstring.ToInt32(s)
			if ok {
				intSlice = append(intSlice, val)
			}
		}
	}

	return intSlice
}

// StringToInt64 将字符串切片转换为int64切片(不可转的跳过)
func StringToInt64(strSlice []string) []int64 {
	var intSlice []int64

	for _, s := range strSlice {
		if cutils.IsNumeric(s) {
			val, ok := cstring.ToInt64(s)
			if ok {
				intSlice = append(intSlice, val)
			}
		}
	}

	return intSlice
}

// IsSlice 检查给定的值是否为array/slice
// 注意它在内部使用了reflect来实现这个特性
func IsSlice(value interface{}) bool {
	rv := reflect.ValueOf(value)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

// IsEmptyWithString 检查切片中是否包含空字符串
func IsEmptyWithString(p []string) bool {
	for _, s := range p {
		if strings.TrimSpace(s) == "" {
			return true
		}
	}
	return false
}
