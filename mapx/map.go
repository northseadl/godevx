package mapx

import "sync"

// Map 包装了一个通用的map类型
type Map[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex // 添加并发安全支持
}

// New 创建一个新的Map实例
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}

// NewWithCapacity 创建一个指定容量的Map实例
func NewWithCapacity[K comparable, V any](capacity int) *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V, capacity),
	}
}

// Keys 返回map的所有键
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values 返回map的所有值
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// HasKeys 检查map是否包含所有指定的键
func HasKeys[K comparable, V any](m map[K]V, ks ...K) bool {
	for _, k := range ks {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

// MapByFunc 通过函数将切片转换为map
func MapByFunc[T any, K comparable, V any](s []T, fn func(item T) (K, V)) map[K]V {
	m := make(map[K]V, len(s))
	for i := range s {
		k, v := fn(s[i])
		m[k] = v
	}
	return m
}

// MapByKey 通过键函数将切片转换为map
func MapByKey[T any, K comparable](s []T, keyFn func(item T) K) map[K]T {
	m := make(map[K]T, len(s))
	for i := range s {
		k := keyFn(s[i])
		m[k] = s[i]
	}
	return m
}

// Merge 合并两个或多个map
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Filter 根据过滤函数筛选map元素
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Transform 转换map的值
func Transform[K comparable, V1 any, V2 any](m map[K]V1, fn func(V1) V2) map[K]V2 {
	result := make(map[K]V2, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// SafeMap方法

func (m *Map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}
