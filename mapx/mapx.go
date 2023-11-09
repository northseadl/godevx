package mapx

// Keys returns the keys of the given map.
//
//	返回给定 map 的键。
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns the values of the given map.
//
//	返回给定 map 的值。
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// HasKeys returns whether the given map has the given keys.
//
//	返回给定 map 是否有给定的键。
func HasKeys[K comparable, V any](m map[K]V, ks ...K) bool {
	for _, k := range ks {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

// MapByFunc returns a map by the given slice and function.
//
//	根据给定的切片和函数返回一个 map。
func MapByFunc[T any, K comparable, V any](s []T, fun func(item T) (K, V)) map[K]V {
	m := make(map[K]V, len(s))
	for i := range s {
		k, v := fun(s[i])
		m[k] = v
	}
	return m
}

// MapByKey returns a map by the given slice and key function.
//
//	根据给定的切片和键函数返回一个 map。
func MapByKey[T any, K comparable](s []T, keyFun func(item T) K) map[K]T {
	m := make(map[K]T, len(s))
	for i := range s {
		k := keyFun(s[i])
		m[k] = s[i]
	}
	return m
}
