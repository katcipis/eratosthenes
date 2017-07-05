package eratosthenes_test

import (
	"testing"

	"github.com/katcipis/eratosthenes"
)

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

type sieveTestCase struct {
	name   string
	n      int
	result []int
}

type sieveFunc func(int) []int

func testSieve(t *testing.T, sieve sieveFunc) {
	cases := []sieveTestCase{
		sieveTestCase{
			name:   "noPrime",
			n:      1,
			result: []int{},
		},
		sieveTestCase{
			name:   "onePrime",
			n:      2,
			result: []int{2},
		},
		sieveTestCase{
			name: "simple",
			n:    30,
			result: []int{
				2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
			},
		},
		sieveTestCase{
			name: "nIsPrime",
			n:    29,
			result: []int{
				2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			result := sieve(testCase.n)
			assertEqualInts(t, testCase.result, result)
		})
	}
}

func BenchmarkSequentialSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eratosthenes.Sequential(100000000)
	}
}
