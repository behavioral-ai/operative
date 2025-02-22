package agent1

import (
	"fmt"
	"math/rand"
)

func newObservation() observation {
	var o observation

	minN := 10
	maxN := 3500
	o.latency = rand.Intn(maxN-minN+1) + minN

	minN = 0
	maxN = 100
	o.gradient = rand.Intn(maxN-minN+1) + minN
	return o

}

func ExampleNewObservation() {
	o := newObservation()
	fmt.Printf("test: NewObservation() -> [lat:%v] [grad:%v]\n", o.Latency(), o.Gradient())

	//o = newObservation()
	//fmt.Printf("test: NewObservation() -> [lat:%v] [grad:%v]\n", o.Latency, o.Gradient)

	//Output:
	//test: NewObservation() -> [lat:962] [grad:42]

}
