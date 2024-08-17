package iterutil_test

import (
	"fmt"

	iterutil "github.com/mazrean/iter-util"
)

func ExampleStack() {
	// Create a new stack
	s := iterutil.NewStack[int]()

	// Push some elements
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for v := range s.Iter {
		fmt.Println(v)
	}

	// Output:
	// 3
	// 2
	// 1
}
