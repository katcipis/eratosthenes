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

	primes := []int{}
	currentPrime := 2
	currentSequence := sequence(currentPrime)

	for currentPrime <= n {
		currentPrime = <-currentSequence
		if currentPrime > n {
			break
		}
		currentSequence = filterNotMultiplesOf(
			currentPrime,
			currentSequence,
		)
		primes = append(primes, currentPrime)
	}

	return primes
}
