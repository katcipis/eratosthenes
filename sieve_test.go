package sieve

import "testing"

func assertEqualInts(t *testing.T, a []int, b []int) {
	if len(a) != len(b) {
		t.Fatal("Not same size: ", a, b)
	}

	for i, v := range a {
		if v != b[i] {
			t.Fatal("Arrays differ", a, b)
		}
	}
}

func TestSequentialSieve(t *testing.T) {
	expectedResult := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	}
	result := sieve(30)
	assertEqualInts(t, expectedResult, result)
}
