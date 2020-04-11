package type_test

import (
	"testing"
)

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(t)
	t.Log(a,aPtr)
	t.Logf("%T,%T,%T",a,aPtr,t)

	var x *int = &a
	t.Log(x,*x,*aPtr)

}

func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*", len(s))
	if s=="" {
		t.Log("nil")
	}
	var r rune = 0x4e2d
	t.Log(string(r),"\u4e2d")
	t.Log('1')

}