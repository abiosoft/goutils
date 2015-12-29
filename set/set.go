// Package set implements a Set using map.
package set

import "sync"

// Set stores distinct items.
// An empty Set struct is not valid for use, use NewSet instead.
//
// Set is safe for concurrent use and all methods can be
// accessed from multiple goroutines.
type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}

// NewSet creates a new Set.
func New() *Set {
	return &Set{make(map[interface{}]struct{}), sync.RWMutex{}}
}

// Add adds a value to the set.
func (s *Set) Add(value interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[value] = struct{}{}
}

// AddAll adds all values to the set distinctly.
func (s *Set) AddAll(values ...interface{}) {
	s.Lock()
	defer s.Unlock()
	for _, value := range values {
		s.m[value] = struct{}{}
	}
}

// AddList adds all items in list to the set distinctlly.
func (s *Set) AddList(list []interface{}) {
	s.AddAll(list...)
}

// Remove removes value from the set if it exists in the set.
func (s *Set) Remove(value interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, value)
}

// RemoveAll removes all values from the set if they exist in the set.
func (s *Set) RemoveAll(values ...interface{}) {
	s.Lock()
	defer s.Unlock()
	for _, value := range values {
		delete(s.m, value)
	}
}

// RemoveList removes all items in list from the set if they exist in
// the set.
func (s *Set) RemoveList(list []interface{}) {
	s.RemoveAll(list...)
}

// Contains check if value exists in the set.
func (s *Set) Contains(value interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[value]
	return ok
}

// ContainsAll checks if all values exist in the set.
func (s *Set) ContainsAll(values ...interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	for _, value := range values {
		_, ok := s.m[value]
		if !ok {
			return false
		}
	}
	return true
}

// ContainsList checks if all items in list exist in the set.
func (s *Set) ContainsList(list []interface{}) bool {
	return s.ContainsAll(list...)
}

// ContainsFunc iterates all the items in the set and passes
// each to f. It returns true the first time a call to f returns
// true and false if no call to f returns true.
func (s *Set) ContainsFunc(f func(interface{}) bool) bool {
	s.RLock()
	defer s.RUnlock()
	for k := range s.m {
		if f(k) {
			return true
		}
	}
	return false
}

// Size returns the number of items in the set.
func (s *Set) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.m)
}

// Clear empties the set.
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[interface{}]struct{})
}

// Iterator returns a new Iterator to iterate through values in the set.
// Each call to this method creates a new Iterator. Therefore, the
// returned Iterator should be assigned to a variable before usage.
//
// It is safe to create and use multiple Iterators in multiple goroutines.
func (s *Set) Iterator() Iterator {
	s.RLock()
	defer s.RUnlock()
	iterChan := make(chan interface{}, len(s.m))
	for k := range s.m {
		iterChan <- k
	}
	close(iterChan)
	return IterFunc(func() (interface{}, bool) {
		value, ok := <-iterChan
		return value, ok
	})
}

// IteratorFunc is same as Iterator but the returned Iterator only
// iterates through values that when passed to f, f returns true.
func (s *Set) IteratorFunc(f func(value interface{}) bool) Iterator {
	s.RLock()
	defer s.RUnlock()
	iterChan := make(chan interface{}, len(s.m))
	for k := range s.m {
		if f(k) {
			iterChan <- k
		}
	}
	close(iterChan)
	return IterFunc(func() (interface{}, bool) {
		value, ok := <-iterChan
		return value, ok
	})
}

// Items returns a slice of all items in the set.
// Modification to the returned slice does not affect the structure of
// the set. However, any item in the set that is a pointer will be affected
// if modified.
func (s *Set) Items() []interface{} {
	s.RLock()
	defer s.RUnlock()
	items := make([]interface{}, len(s.m))
	i := 0
	for k := range s.m {
		items[i] = k
		i++
	}
	return items
}

// ItemsFunc returns slice of all items that when passed to f, f returns true.
// Modification to the returned slice does not affect the structure of
// the set. However, any item in the set that is a pointer will be affected
// if modified.
func (s *Set) ItemsFunc(f func(value interface{}) bool) []interface{} {
	s.RLock()
	defer s.RUnlock()
	items := []interface{}{}
	for k := range s.m {
		if f(k) {
			items = append(items, k)
		}
	}
	return items
}

// Iterator iterates through a group of items.
type Iterator interface {
	// HasNext checks if there is a next value and moves to it.
	HasNext() bool
	// Value returns the current item. The initial value is nil and requires a call
	// to HasNext before usage. If HasNext returns false, it returns nil.
	Value() interface{}
}

type iterable struct {
	value interface{}
	next  func() (interface{}, bool)
}

func (i *iterable) HasNext() bool {
	value, ok := i.next()
	i.value = value
	return ok
}

func (i *iterable) Value() interface{} {
	return i.value
}

// IterFunc creates an Iterator using f
func IterFunc(f func() (interface{}, bool)) Iterator {
	return &iterable{next: f}
}
