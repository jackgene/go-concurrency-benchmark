package counter

import "testing"

func benchmarkCounter(counter Counter, count int, concurrency int) {
	done := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			for c := 0; c < count; c++ {
				counter.Increment()
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < concurrency; i++ {
		<-done
	}
}

func BenchmarkActorCounter(b *testing.B) {
	counter := NewActorCounter()

	benchmarkCounter(counter, b.N, 10)
}

func BenchmarkMutexCounter(b *testing.B) {
	counter := NewMutexCounter()

	benchmarkCounter(counter, b.N, 10)
}

func BenchmarkUnsafeCounter(b *testing.B) {
	counter := NewUnsafeCounter()

	benchmarkCounter(counter, b.N, 10)
}
