// Package cherrySlice code from: https://github.com/beego/beego/blob/develop/core/utils/slice.go
package cherrySlice

import (
	"math/rand"
	"reflect"
	"strings"

	cstring "github.com/cherry-game/cherry/extend/string"
	cutils "github.com/cherry-game/cherry/extend/utils"
)

type (
	// Addable 定义一个接口，表示支持 + 操作的类型
	Addable interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~complex64 | ~complex128
	}

	// Ordered 定义一个接口，表示支持比较操作的类型
	Ordered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}
)

// CreateWithValue 创建一个指定长度的切片，并初始化所有元素为指定值
func CreateWithValue[V any](v V, length int) []V {
	sl := make([]V, length)
	for i := range sl {
		sl[i] = v
	}
	return sl
}

// IndexOf 返回元素在切片中的位置(-1表示没有)
func IndexOf[V comparable](sl []V, v V) int {
	for i, vv := range sl {
		if vv == v {
			return i
		}
	}
	return -1
}

// CheckIndex 检查索引是否在切片范围内
func CheckIndex[V any](sl []V, index int) bool {
	if index < 0 || index >= len(sl) {
		return false
	}
	return true
}

// Contains 检测切片中是否包含指定值
func Contains[V comparable](sl []V, v V) bool {
	return IndexOf(sl, v) >= 0
}

// SafeSub 安全取子切片
func SafeSub[V any](slice []V, start, end int) []V {
	sliceLen := len(slice)
	if start > sliceLen {
		start = sliceLen
	}
	if end > sliceLen {
		end = sliceLen
	}
	return slice[start:end]
}

// AppendUnique 添加一个值到切片中，如果切片中已经存在该值，则返回false
func AppendUnique[V comparable](sl []V, v V) []V {
	if !Contains(sl, v) {
		sl = append(sl, v)
	}
	return sl
}

// InsertSlice 在切片中插入一个值
func InsertSlice[V any](slice []V, index int, value V) []V {
	if index < 0 || index > len(slice) {
		panic("index out of range")
	}
	return append(slice[:index], append([]V{value}, slice[index:]...)...)
}

// RemoveIndex 删除切片中的指定索引
func RemoveIndex[V any](sl []V, index int) []V {
	if index < 0 || index >= len(sl) {
		return sl
	}

	result := make([]V, 0, len(sl)-1)
	result = append(result, sl[:index]...)
	result = append(result, sl[index+1:]...)
	return result
}

// Remove 删除切片中的指定值
func Remove[V comparable](sl []V, value V) []V {
	result := make([]V, 0, len(sl))
	for _, v := range sl {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// Min 获取切片中的最小值
func Min[V Ordered](slice []V) (min V) {
	if len(slice) == 0 {
		return
	}

	min = slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return
}

// Max 获取切片中的最大值
func Max[V Ordered](slice []V) (max V) {
	if len(slice) == 0 {
		return
	}

	max = slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return
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
func Merge[V any](slice1, slice2 []V) []V {
	result := make([]V, 0, len(slice1)+len(slice2))
	result = append(result, slice1...)
	result = append(result, slice2...)
	return result
}

// Map 生成转换后的新切片
func Map[V1 any, V2 any](slice []V1, a func(V1) V2) (destSlice []V2) {
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

// RandIndex 随机返回切片的索引(如果切片为空，返回-1)
func RandIndex[V any](a []V) int {
	length := len(a)
	if length == 0 {
		return -1
	}

	return rand.Intn(len(a))
}

// Sum 对切片所有元素求和
func Sum[V Addable](intSlice []V) (sum V) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

// Count 统计切片中值的数量
func Count[V comparable](sl []V, v V) (c int) {
	for _, vv := range sl {
		if vv == v {
			c++
		}
	}
	return
}

// CountFunc 统计切片中满足条件的数量
func CountFunc[V any](slice []V, a func(V) bool) (c int) {
	for _, v := range slice {
		if a(v) {
			c++
		}
	}
	return
}

// Find 获取切片中满足条件的首值
func Find[V comparable](sl []V, a func(V) bool) (V, bool) {
	for _, vv := range sl {
		if a(vv) {
			return vv, true
		}
	}

	var zero V
	return zero, false
}

// FindIndex 获取切片中满足条件的首值的索引
func FindIndex[V comparable](sl []V, a func(V) bool) int {
	for i, vv := range sl {
		if a(vv) {
			return i
		}
	}
	return -1
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
		if !Contains(slice2, v) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// Intersect 求切片1和切片2的交集
func Intersect[V comparable](slice1, slice2 []V) (sameSlice []V) {
	for _, v := range slice1 {
		if Contains(slice2, v) {
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

// Range 生成一个从start到end范围指定步长的新切片
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

// ConvertAnySlice 将切片转换为any切片
func ConvertAnySlice[T any](s []T) []any {
	result := make([]any, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
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
