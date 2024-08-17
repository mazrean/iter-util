package iterutil_test

import (
	"fmt"

	iterutil "github.com/mazrean/iter-util"
)

func ExamplePriorityQueue() {
	// Create a new priority queue
	pq := iterutil.NewPriorityQueue[int]()

	// Push some elements
	pq.Push(1)
	pq.Push(3)
	pq.Push(2)

	for v := range pq.Iter {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}
