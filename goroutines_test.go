// goroutines_test
package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	s1 := sumGoRoutines(slice)
	s2 := sumGoPlain(slice)
	if s1 != 4999999950000000 || s2 != 4999999950000000 || s1 != s2 {
		t.Errorf("get result %d \n", s1)
		t.Errorf("get result %d \n", s2)
	}
}
