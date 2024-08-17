package iterutil_test

import (
	"fmt"

	iterutil "github.com/mazrean/iter-util"
)

func ExampleQueue() {
	// Create a new queue
	q := iterutil.NewQueue[int]()

	// Push some elements
	q.Push(1)
	q.Push(2)
	q.Push(3)

	for v := range q.Iter {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}
