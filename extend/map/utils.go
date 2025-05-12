package cherryMap

// HasKey 判断键是否存在
func HasKey[K comparable, V any](m map[K]V, k K) bool {
	_, b := m[k]
	return b
}

// Keys 获取键切片
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m)) // 预分配内存
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取值切片
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
