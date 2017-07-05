package eratosthenes_test

import (
	"testing"

	"github.com/katcipis/eratosthenes"
)

func TestConcurrentSieve(t *testing.T) {
	testSieve(t, eratosthenes.Concurrent)
}
