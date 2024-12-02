package setx

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

// Set 定义集合的基本接口
type Set[T comparable] interface {
	Add(items ...T)
	Remove(items ...T)
	Contains(item T) bool
	Len() int
	Clear()
	IsEmpty() bool

	// 集合操作
	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]
	SymmetricDifference(other Set[T]) Set[T]
	IsSubset(other Set[T]) bool
	IsSuperset(other Set[T]) bool
	Equal(other Set[T]) bool

	// 转换操作
	ToSlice() []T
	Clone() Set[T]
	String() string

	// 功能性操作
	ForEach(f func(T))
	Filter(f func(T) bool) Set[T]
	Map(f func(T) T) Set[T]
	Any(f func(T) bool) bool
	All(f func(T) bool) bool

	// 迭代器
	Iterator() Iterator[T]
}

// Iterator 定义迭代器接口
type Iterator[T comparable] interface {
	HasNext() bool
	Next() T
	Remove()
}

// HashSet 实现 Set 接口的基础结构
type HashSet[T comparable] struct {
	items map[T]struct{}
}

// ThreadSafeHashSet 线程安全的 HashSet 实现
type ThreadSafeHashSet[T comparable] struct {
	*HashSet[T]
	mu sync.RWMutex
}

// NewHashSet 创建一个新的HashSet实例
func NewHashSet[T comparable](items ...T) *HashSet[T] {
	hs := &HashSet[T]{
		items: make(map[T]struct{}, len(items)),
	}
	hs.Add(items...)
	return hs
}

// NewThreadSafeHashSet 创建一个新的线程安全HashSet实例
func NewThreadSafeHashSet[T comparable](items ...T) *ThreadSafeHashSet[T] {
	return &ThreadSafeHashSet[T]{
		HashSet: NewHashSet(items...),
	}
}

// Add 添加元素到集合
func (hs *HashSet[T]) Add(items ...T) {
	for _, item := range items {
		hs.items[item] = struct{}{}
	}
}

// Remove 从集合中移除元素
func (hs *HashSet[T]) Remove(items ...T) {
	for _, item := range items {
		delete(hs.items, item)
	}
}

// Contains 检查元素是否在集合中
func (hs *HashSet[T]) Contains(item T) bool {
	_, exists := hs.items[item]
	return exists
}

// Len 返回集合中元素的数量
func (hs *HashSet[T]) Len() int {
	return len(hs.items)
}

// Clear 清空集合
func (hs *HashSet[T]) Clear() {
	hs.items = make(map[T]struct{})
}

// IsEmpty 检查集合是否为空
func (hs *HashSet[T]) IsEmpty() bool {
	return len(hs.items) == 0
}

// Union 求两个集合的并集
func (hs *HashSet[T]) Union(other Set[T]) Set[T] {
	result := NewHashSet[T]()
	for item := range hs.items {
		result.Add(item)
	}
	for _, item := range other.ToSlice() {
		result.Add(item)
	}
	return result
}

// Intersection 求两个集合的交集
func (hs *HashSet[T]) Intersection(other Set[T]) Set[T] {
	result := NewHashSet[T]()
	for item := range hs.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference 求两个集合的差集
func (hs *HashSet[T]) Difference(other Set[T]) Set[T] {
	result := NewHashSet[T]()
	for item := range hs.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// SymmetricDifference 求两个集合的对称差
func (hs *HashSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := hs.Union(other)
	intersection := hs.Intersection(other)
	return result.Difference(intersection)
}

// IsSubset 检查是否为其他集合的子集
func (hs *HashSet[T]) IsSubset(other Set[T]) bool {
	for item := range hs.items {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSuperset 检查是否为其他集合的超集
func (hs *HashSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(hs)
}

// Equal 检查两个集合是否相等
func (hs *HashSet[T]) Equal(other Set[T]) bool {
	if hs.Len() != other.Len() {
		return false
	}
	return hs.IsSubset(other)
}

// ToSlice 将集合转换为切片
func (hs *HashSet[T]) ToSlice() []T {
	result := make([]T, 0, len(hs.items))
	for item := range hs.items {
		result = append(result, item)
	}
	return result
}

// Clone 克隆集合
func (hs *HashSet[T]) Clone() Set[T] {
	clone := NewHashSet[T]()
	for item := range hs.items {
		clone.Add(item)
	}
	return clone
}

// String 返回集合的字符串表示
func (hs *HashSet[T]) String() string {
	items := hs.ToSlice()
	strItems := make([]string, len(items))
	for i, item := range items {
		strItems[i] = fmt.Sprintf("%v", item)
	}
	sort.Strings(strItems) // 保证输出顺序一致
	return fmt.Sprintf("Set{%s}", strings.Join(strItems, ", "))
}

// ForEach 对集合中的每个元素执行操作
func (hs *HashSet[T]) ForEach(f func(T)) {
	for item := range hs.items {
		f(item)
	}
}

// Filter 过滤集合中的元素
func (hs *HashSet[T]) Filter(f func(T) bool) Set[T] {
	result := NewHashSet[T]()
	for item := range hs.items {
		if f(item) {
			result.Add(item)
		}
	}
	return result
}

// Map 对集合中的元素进行映射
func (hs *HashSet[T]) Map(f func(T) T) Set[T] {
	result := NewHashSet[T]()
	for item := range hs.items {
		result.Add(f(item))
	}
	return result
}

// Any 检查是否存在满足条件的元素
func (hs *HashSet[T]) Any(f func(T) bool) bool {
	for item := range hs.items {
		if f(item) {
			return true
		}
	}
	return false
}

// All 检查是否所有元素都满足条件
func (hs *HashSet[T]) All(f func(T) bool) bool {
	for item := range hs.items {
		if !f(item) {
			return false
		}
	}
	return true
}

// MarshalJSON 实现JSON序列化
func (hs *HashSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(hs.ToSlice())
}

// UnmarshalJSON 实现JSON反序列化
func (hs *HashSet[T]) UnmarshalJSON(data []byte) error {
	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	hs.Clear()
	hs.Add(items...)
	return nil
}

// ThreadSafeHashSet的方法实现
func (ts *ThreadSafeHashSet[T]) Add(items ...T) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.HashSet.Add(items...)
}

func (ts *ThreadSafeHashSet[T]) Remove(items ...T) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.HashSet.Remove(items...)
}

func (ts *ThreadSafeHashSet[T]) Contains(item T) bool {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.HashSet.Contains(item)
}

func (ts *ThreadSafeHashSet[T]) Clear() {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.HashSet.Clear()
}

func (ts *ThreadSafeHashSet[T]) ToSlice() []T {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.HashSet.ToSlice()
}

func (ts *ThreadSafeHashSet[T]) ForEach(f func(T)) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	ts.HashSet.ForEach(f)
}

func (ts *ThreadSafeHashSet[T]) Union(other Set[T]) Set[T] {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.HashSet.Union(other)
}

func (ts *ThreadSafeHashSet[T]) Intersection(other Set[T]) Set[T] {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.HashSet.Intersection(other)
}

// 迭代器实现
type hashSetIterator[T comparable] struct {
	set  *HashSet[T]
	keys []T
	pos  int
}

func (hs *HashSet[T]) Iterator() Iterator[T] {
	return &hashSetIterator[T]{
		set:  hs,
		keys: hs.ToSlice(),
		pos:  0,
	}
}

func (it *hashSetIterator[T]) HasNext() bool {
	return it.pos < len(it.keys)
}

func (it *hashSetIterator[T]) Next() T {
	if !it.HasNext() {
		panic("No more elements")
	}
	item := it.keys[it.pos]
	it.pos++
	return item
}

func (it *hashSetIterator[T]) Remove() {
	if it.pos <= 0 {
		panic("Invalid state for removal")
	}
	it.set.Remove(it.keys[it.pos-1])
}
