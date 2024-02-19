package slicex

// Pluck returns a new slice by the given slice and function.
//
//	根据给定的切片和函数返回一个新的切片。
func Pluck[T any, V comparable](s []T, getValue func(item T) V) []V {
	var ks []V
	for _, t := range s {
		ks = append(ks, getValue(t))
	}
	return ks
}

// Contains returns whether the given slice contains the given values.
//
//	返回给定的切片是否包含给定的值。
func Contains[T comparable](slice []T, values ...T) bool {
	var set map[T]struct{}
	set = make(map[T]struct{}, 0)
	for _, i := range slice {
		set[i] = struct{}{}
	}
	for _, value := range values {
		if _, ok := set[value]; !ok {
			return false
		}
	}
	return true
}

// Filter returns a filtered slice by the given slice and function.
//
//	根据给定的切片和函数返回一个过滤后的切片。
func Filter[T any](slice []T, fun func(item T) bool) []T {
	var s []T
	for _, item := range slice {
		if fun(item) {
			s = append(s, item)
		}
	}
	return s
}

// Unique returns unique slice by the given slice.
//
//	根据给定的切片返回一个去重的切片。
func Unique[T comparable](s []T) []T {
	var result []T
	set := make(map[T]struct{})
	for _, item := range s {
		if _, ok := set[item]; !ok {
			set[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Intersect returns the intersection of the given slices.
//
//	返回给定切片的交集。
func Intersect[T comparable](s1, s2 []T) []T {
	var result []T
	set := make(map[T]struct{})
	for _, item := range s1 {
		set[item] = struct{}{}
	}
	for _, item := range s2 {
		if _, ok := set[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

// Difference returns the difference of the given slices.
//
//	返回给定切片的差集。
func Difference[T comparable](s1, s2 []T) []T {
	var result []T
	set := make(map[T]struct{})
	for _, item := range s1 {
		set[item] = struct{}{}
	}
	for _, item := range s2 {
		if _, ok := set[item]; ok {
			delete(set, item)
		}
	}
	for item := range set {
		result = append(result, item)
	}
	return result
}

// Union returns the union of the given slices.
//
//	返回给定切片的并集。
func Union[T comparable](s1, s2 []T) []T {
	var result []T
	set := make(map[T]struct{})
	for _, item := range s1 {
		set[item] = struct{}{}
	}
	for _, item := range s2 {
		set[item] = struct{}{}
	}
	for item := range set {
		result = append(result, item)
	}
	return result
}
