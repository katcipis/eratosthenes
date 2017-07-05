package eratosthenes_test

import (
	"testing"

	"github.com/katcipis/eratosthenes"
)

func TestSequentialSieve(t *testing.T) {
	testSieve(t, eratosthenes.Sequential)
}
