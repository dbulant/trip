// goroutines_test
package main

import (
	"fmt"
	"os"
	"testing"
)

var slice []uint64

func BenchmarkBigSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumGoRoutines(slice)
	}
}

func BenchmarkBigSumPlain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumGoPlain(slice)
	}
}

func TestMain(m *testing.M) {
	slice = generateLargeSlice(bigNumber)
	res1 := testing.Benchmark(BenchmarkBigSum)
	res2 := testing.Benchmark(BenchmarkBigSumPlain)
	//res1 is  10  104285020 ns/op test run 10 times, one loop took 104285020 ns
	fmt.Printf("res1 is %+v \n", res1)
	//res2 is  10  132683841 ns/op test run 10 times, one loop took 132683841 ns
	fmt.Printf("res2 is %+v \n", res2)
	os.Exit(m.Run())
}
