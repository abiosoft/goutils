package set

import (
	"fmt"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	s := NewSet()
	s.Add("a")
	if s.Size() != 1 {
		t.Errorf("Expected 1 found %v", s.Size())
	}
	s.Add("a")
	if s.Size() != 1 {
		t.Errorf("Expected 1 found %v", s.Size())
	}
	if _, ok := s.m["a"]; !ok {
		t.Errorf("a should be present in the set")
	}
	s.AddAll("b", "c", "d")
	if s.Size() != 4 {
		t.Errorf("Expected 4 found %v", s.Size())
	}
	for _, v := range []interface{}{"b", "c", "d"} {
		if _, ok := s.m[v]; !ok {
			t.Error(v, " should be present in the set")
		}
	}
	s.AddList([]interface{}{"e", "f", "g"})
	if s.Size() != 7 {
		t.Errorf("Expected 7 found %v", s.Size())
	}
	for _, v := range []interface{}{"e", "f", "g"} {
		if _, ok := s.m[v]; !ok {
			t.Error(v, " should be present in the set")
		}
	}
	s = NewSet()
	s.Add(1)
	s.Add("1")
	if s.Size() != 2 {
		t.Errorf("Expected 2 found %v", 2)
	}
	s = NewSet()
	var wg sync.WaitGroup
	add := func(start int) {
		for i := start; i < start+10; i++ {
			s.Add(i)
		}
		wg.Done()
	}
	wg.Add(3)
	go add(0)
	go add(10)
	go add(20)
	wg.Wait()

	if s.Size() != 30 {
		t.Errorf("Expected 30 found %v", s.Size())
	}

	for i := 0; i < 30; i++ {
		if _, ok := s.m[i]; !ok {
			t.Error(i, "should be present in the set")
		}
	}
}

func TestRemove(t *testing.T) {
	s := NewSet()
	s.AddAll("a", "b")
	s.Remove("a")
	if s.Size() != 1 {
		t.Error("Expected 1 found %v", s.Size())
	}

	s.Remove("b")
	if s.Size() != 0 {
		t.Error("Expected 0 found %v", s.Size())
	}

	s.AddAll(1, 2, 2, 3, 4, 1, 1, 99, 0)
	s.RemoveAll(0, 1, 2, 3, 4, 99)
	if s.Size() != 0 {
		t.Error("Expected 0 found %v", s.Size())
	}

	s.AddAll(1, 2, 2, 3, 4, 1, 1, 99, 0)
	s.RemoveList([]interface{}{0, 1, 2, 3, 4, 99})
	if s.Size() != 0 {
		t.Error("Expected 0 found %v", s.Size())
	}

	s.AddAll(1, 2, 2, 3, 4, 1, 1, 99, 0)
	s.RemoveList([]interface{}{0, 1, 99})
	if s.Size() != 3 {
		t.Error("Expected 3 found %v", s.Size())
	}

	s.Clear()
	if s.Size() != 0 {
		t.Error("Expected 0 found %v", s.Size())
	}

	s.AddAll(1, 2, 10, 11, 772367, 1, 2, 3, 4, 1, 1, 99, 0)
	s.Clear()
	if s.Size() != 0 {
		t.Error("Expected 3 found %v", s.Size())
	}

}

func TestContains(t *testing.T) {
	s := NewSet()
	s.Add("a")
	if !s.Contains("a") {
		t.Errorf("a should be present in the set")
	}
	if s.Contains(1) {
		t.Errorf("1 should not be presen in the set")
	}
	s.AddAll(2, 3, 9, 0)
	if s.Size() != 5 {
		t.Errorf("Expected 5 found %v", s.Size())
	}
	if !s.ContainsAll(0, 3, 2) {
		t.Errorf("0,3,2 should be present in the set")
	}
	if s.ContainsAll(4, 5, 6) {
		t.Errorf("4,5,6 should not be present in the set")
	}
	if !s.ContainsList([]interface{}{2, 3, 9}) {
		t.Errorf("2,3,9 must be present in the set")
	}
	if !s.ContainsFunc(func(v interface{}) bool {
		_, ok := v.(string)
		return ok
	}) {
		t.Errorf("a string should be present in the set")
	}
	if s.ContainsFunc(func(v interface{}) bool {
		_, ok := v.(bool)
		return ok
	}) {
		t.Errorf("a bool should not be present in the set")
	}
}

func TestItems(t *testing.T) {
	s := NewSet()
	s.AddAll(2, 3, 9, 0)
	items := s.Items()

	for _, v := range items {
		if v != 2 && v != 3 && v != 9 && v != 0 {
			t.Error(v, "should be present in items")
		}
	}

	items = s.ItemsFunc(func(v interface{}) bool {
		return v.(int) > 2
	})
	for _, v := range items {
		if v != 3 && v != 9 {
			t.Error(v, "should be present in items")
		}
	}
}

func TestIteration(t *testing.T) {
	s := NewSet()

	s.AddAll(1, 2, 3, 4)

	iter := s.Iterator()
	for iter.HasNext() {
		v := iter.Value()
		if v != 1 && v != 2 && v != 3 && v != 4 {
			t.Error(v, "should be part of interated values")
		}
	}

	iter = s.IteratorFunc(func(v interface{}) bool {
		return v.(int) > 2
	})
	for iter.HasNext() {
		v := iter.Value()
		if v != 3 && v != 4 {
			t.Error(v, "should be part of interated values")
		}
	}

	iter = s.Iterator()
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		s.Add(5)
		w.Done()
	}()
	for iter.HasNext() {
		if s.Contains(5) {
			t.Errorf("5 should not be present yet. Writers should be blocked until iteration is done")
		}
	}
	// wait until addition is done
	w.Wait()
	if !s.Contains(5) {
		t.Errorf("5 should be present now. Iteration is done")
	}
}

func Example() {
	s := NewSet()

	// stores distinctly, repetitions have no effect.
	s.AddAll(1, 1, 2, 2, 3, 6, 2, 1, "a", "b", "a")

	fmt.Println(s.Size()) // 6

	list := []interface{}{1, 2, 3, 6, "a", "b"}
	fmt.Println(s.ContainsList(list)) // true

	iter := s.Iterator()
	for iter.HasNext() {
		// ... do something with iter.Value()
		fmt.Println("iterating")
	}

	// Output: 6
	// true
	// iterating
	// iterating
	// iterating
	// iterating
	// iterating
	// iterating
}

func ExampleIterator() {
	s := NewSet()

	s.AddAll(1, 1, 1)

	iter := s.Iterator()

	for iter.HasNext() {
		fmt.Println(iter.Value())
	}

	// Output: 1
}

func ExampleSet_IteratorFunc() {
	s := NewSet()

	s.AddAll(1, 2, 3, 4)

	iter := s.IteratorFunc(func(value interface{}) bool {
		return value.(int) > 3
	})

	for iter.HasNext() {
		fmt.Println(iter.Value())
	}

	// Output: 4
}
