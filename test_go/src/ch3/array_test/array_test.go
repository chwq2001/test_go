package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	var arr1 [3]int
	arr2 := [...] int {1,2,3}
	arr3 := [] int {1,2,3}
	t.Logf("%T,%T,%T",arr1,arr2,arr3)
	for i:=0; i<len(arr1); i++ {
		t.Log(arr1[i])
	}

	for i:=0; i<len(arr2); i++ {
		t.Log(arr2[i])
	}
}

func TestArrayTravel(t *testing.T) {
	arr := [][] int {{1,2,3},{3,2,1},{2,3,1}}
	for i,r := range arr {
		for j,c := range r {
			t.Log(i,j,c)
		}
	}
}

func TestArraySection(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	t.Log(a[2:3])
	t.Log(a[2:4])
	t.Log(a[2:])
	t.Log(a[:4])
	t.Log(a[2:len(a)])
	t.Log(a[:])

	t.Log(len(a), cap(a))
	//printArray(a,t)
}