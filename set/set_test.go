package set

import "fmt"

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
