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

type sieveTestCase struct {
	name   string
	n      int
	result []int
}

func TestSequentialSieve(t *testing.T) {
	cases := []sieveTestCase{
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
		sieve(100000000)
	}
}
