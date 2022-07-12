package counter

import "testing"

func incrementCounter(counter Counter, count int, concurrency int) {
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

func TestCounter(t *testing.T) {
	t.Run("Serially increment ActorCounter", func(t *testing.T) {
		// Set up
		counter := NewActorCounter()

		// Test
		incrementCounter(counter, 100_000, 1)

		// Verify
		var want uint = 100_000
		got := counter.Get()
		if want != got {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Serially increment MutexCounter", func(t *testing.T) {
		// Set up
		counter := NewMutexCounter()

		// Test
		incrementCounter(counter, 100_000, 1)

		// Verify
		var want uint = 100_000
		got := counter.Get()
		if want != got {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Serially increment UnsafeCounter", func(t *testing.T) {
		// Set up
		counter := NewUnsafeCounter()

		// Test
		incrementCounter(counter, 100_000, 1)

		// Verify
		var want uint = 100_000
		got := counter.Get()
		if want != got {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Concurrently increment ActorCounter", func(t *testing.T) {
		// Set up
		counter := NewActorCounter()

		// Test
		incrementCounter(counter, 100_000, 2)

		// Verify
		var want uint = 200_000
		got := counter.Get()
		if want != got {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Concurrently increment MutexCounter", func(t *testing.T) {
		// Set up
		counter := NewMutexCounter()

		// Test
		incrementCounter(counter, 100_000, 2)

		// Verify
		var want uint = 200_000
		got := counter.Get()
		if want != got {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Concurrently increment UnsafeCounter", func(t *testing.T) {
		// Set up
		counter := NewUnsafeCounter()

		// Test
		incrementCounter(counter, 100_000, 2)

		// Verify
		// This one is expected to fail
		//var want uint = 200_000
		//got := counter.Get()
		//if want != got {
		//	t.Errorf("wanted %d, got %d", want, got)
		//}
	})
}

func BenchmarkCounter(b *testing.B) {
	b.Run("Benchmark ActorCounter", func(b *testing.B) {
		counter := NewActorCounter()

		incrementCounter(counter, b.N, 10)
	})

	b.Run("Benchmark MutexCounter", func(b *testing.B) {
		counter := NewMutexCounter()

		incrementCounter(counter, b.N, 10)
	})

	b.Run("Benchmark UnsafeCounter", func(b *testing.B) {
		counter := NewUnsafeCounter()

		incrementCounter(counter, b.N, 10)
	})
}
