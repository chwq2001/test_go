package fib_test

import (
	"fmt"
	"math"
	"testing"
)

func TestFibList(t *testing.T)  {
	var (
		a int = 1
		b int = 1
	)
	t.Log(a)
	var i int  = 2
	for ; i< 11+2; i++ {
		t.Log(" ",b)
		//var tmp = b
		//b = a + b
		//a = tmp
		a,b = b,a+b
	}
	fmt.Println()
	var sqrt5 float64 = math.Sqrt(5)
	nv := (math.Pow((sqrt5+1)/2, float64(i-1)) - (math.Pow((1-sqrt5)/2, float64(i-1)))) / sqrt5
	t.Log(math.Round(nv))
}