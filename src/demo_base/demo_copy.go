package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	fmt.Println(scores)
	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst) //[0 2 2 15 26]

	copy(worst[2:4], scores[:5])

	fmt.Println(worst) //[0 2 0 2 26]
}
