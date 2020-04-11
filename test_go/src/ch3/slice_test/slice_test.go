package slice_test

import "testing"

func TestSliceSection(t *testing.T) {
	x := []int {1,2,3}

	t.Log(len(x), cap(x))
	x = append(x, 1)
	t.Log(len(x), cap(x))

	var a []int
	t.Log(len(a), cap(a))
	b := append(a, 5)
	t.Log(len(a), cap(a))
	t.Log(len(b), cap(b))

	c := []int{1, 2}
	t.Log(len(c), cap(c))
	/*
		make(t Type, len,capacity)
	*/
	d := make([]int, 3, 5)
	t.Logf("%T,%d,%d", d, len(d), cap(d))
	t.Log(d[0], d[1], d[2])
	e := append(d, '6')
	t.Logf("d: %T,%d,%d", d, len(d), cap(d))
	t.Logf("e: %T,%d,%d", e, len(e), cap(e))

	f := append(append(append(d, '6'), 7), 8)
	t.Logf("d: %T,%d,%d", d, len(d), cap(d))
	t.Logf("f: %T,%d,%d", e, len(f), cap(f))
}
