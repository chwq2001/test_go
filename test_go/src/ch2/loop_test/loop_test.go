package loop_test

import (
	"math"
	"testing"
)

func TestLoopFor(t *testing.T) {

	for n := 0; n < 5; n++ {
		t.Log(n)
	}
	t.Log("for end")
}

func TestLoopWhile(t *testing.T) {
	n := 0
	for n<=5{
		n++
		t.Log(n)
	}
	t.Log("while  end")
}

func TestLoopWhileTrue(t *testing.T) {
	n := 0
	for {
		n++
		t.Log(n)
		if n > 5 {
			t.Log("WhileTrue end")
			break
		}
	}
}

func TestLoopContinue(t *testing.T) {

	for n := 0; n < 10; n++ {
		if n%2!=0 {
			continue
		} else if n%3==0 { //不能换行
			t.Log(math.Pow(float64(n),2))
		} else { //不能换行
			t.Log(n)
		}

	}
	t.Log("Continue  end")
}