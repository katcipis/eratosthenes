package eratosthenes

func nextPrime(n int, prime int, possiblePrimes []bool) int {
	for i := prime; i <= n; i += prime {
		possiblePrimes[i] = true
	}
	for i := prime + 1; i <= n; i++ {
		if possiblePrimes[i] == false {
			return i
		}
	}
	return n + 1
}

func Sequential(n int) []int {
	possiblePrimes := make([]bool, n+1)
	prime := 2
	results := []int{}
	for prime <= n {
		results = append(results, prime)
		prime = nextPrime(n, prime, possiblePrimes)
	}

	return results
}
