package main

import "go-concurrency-performance/counter"

const concurrency = 20

func incrementAll(counter counter.Counter, completion chan struct{}) {
	for i := 0; i < 1_000_000; i++ {
		counter.Increment()
	}
	completion <- struct{}{}
}

func main() {
	//c := counter.NewUnsafeCounter()
	c := counter.NewActorCounter()
	//c := counter.NewMutexCounter()

	comp := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		go incrementAll(c, comp)
	}

	for i := 0; i < concurrency; i++ {
		<-comp
	}
	println(c.Get())
}
