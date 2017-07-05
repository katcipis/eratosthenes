package eratosthenes

func sequence(i int) <-chan int {
	seq := make(chan int)

	go func() {
		// This will obviously leave a locked goroutine
		for {
			seq <- i
			i += 1
		}
	}()

	return seq
}

func filterNotMultiplesOf(n int, seq <-chan int) <-chan int {
	filtered := make(chan int)

	go func() {
		for i := range seq {
			if i%n != 0 {
				filtered <- i
			}
		}
	}()

	return filtered
}

func Concurrent(n int) []int {
	// This algorithm will create O(N) goroutines and channels
	// It will be a chain of filter channels fed by the first sequence generator

	primes := []int{}
	prime := 2
	primegenerator := sequence(prime)

	for {
		prime = <-primegenerator
		if prime > n {
			break
		}
		primegenerator = filterNotMultiplesOf(
			prime,
			primegenerator,
		)
		primes = append(primes, prime)
	}

	return primes
}
