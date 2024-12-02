package slicex

import "github.com/northseadl/godevx/setx"

// Pluck 根据给定的切片和函数返回一个新的切片
func Pluck[T any, V comparable](s []T, getValue func(item T) V) []V {
	result := make([]V, 0, len(s))
	for _, t := range s {
		result = append(result, getValue(t))
	}
	return result
}

// Contains 返回给定的切片是否包含给定的值
func Contains[T comparable](slice []T, values ...T) bool {
	set := setx.NewHashSet(slice...)
	for _, value := range values {
		if !set.Contains(value) {
			return false
		}
	}
	return true
}

// Filter 根据给定的切片和函数返回一个过滤后的切片
func Filter[T any](slice []T, fun func(item T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if fun(item) {
			result = append(result, item)
		}
	}
	return result
}

// Unique 根据给定的切片返回一个去重的切片
func Unique[T comparable](s []T) []T {
	set := setx.NewHashSet(s...)
	return set.ToSlice()
}

// Intersect 返回给定切片的交集
func Intersect[T comparable](s1, s2 []T) []T {
	set1 := setx.NewHashSet(s1...)
	set2 := setx.NewHashSet(s2...)
	return set1.Intersection(set2).ToSlice()
}

// Difference 返回给定切片的差集
func Difference[T comparable](s1, s2 []T) []T {
	set1 := setx.NewHashSet(s1...)
	set2 := setx.NewHashSet(s2...)
	return set1.Difference(set2).ToSlice()
}

// Union 返回给定切片的并集
func Union[T comparable](s1, s2 []T) []T {
	set1 := setx.NewHashSet(s1...)
	set2 := setx.NewHashSet(s2...)
	return set1.Union(set2).ToSlice()
}
