package intsqrt

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSqrt32(t *testing.T) {
	var tests = []struct {
		n, r uint32
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{8, 2},
		{9, 3},
		{15, 3},
		{16, 4},
		{0xffffffff, 0xffff},
	}
	for _, test := range tests {
		r := Sqrt32(test.n)
		if test.r != r {
			t.Errorf("for n=%v got %v, want %v", test.n, r, test.r)
		}
	}
}

func BenchmarkSqrt32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Sqrt32(0x1000000)
		if r != 0x1000 {
			fmt.Println("wrong result")
		}
	}
}

func TestSqrt64(t *testing.T) {
	var tests = []struct {
		n, r uint64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{8, 2},
		{9, 3},
		{15, 3},
		{16, 4},
		{0xffffffff, 0xffff},
		{0xffffffff*0xffffffff - 1, 0xfffffffe},
		{0xffffffff * 0xffffffff, 0xffffffff},
		{0xffffffffffffffff, 0xffffffff},
	}
	for _, test := range tests {
		r := Sqrt64(test.n)
		if test.r != r {
			t.Errorf("for n=%v got %v, want %v", test.n, r, test.r)
		}
	}
}

func BenchmarkSqrt64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Sqrt64(40000000000000000)
		if r != 200000000 {
			fmt.Println("wrong result")
		}
	}
}

func BenchmarkSqrtBig(b *testing.B) {
	v := new(big.Int).SetUint64(40000000000000000)
	exp := big.NewInt(200000000)
	for n := 0; n < b.N; n++ {
		r := new(big.Int).Sqrt(v)
		if r.Cmp(exp) != 0 {
			fmt.Println("wrong result")
		}
	}
}
