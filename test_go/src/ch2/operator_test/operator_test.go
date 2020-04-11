package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...] int {1,2,3,4}
	b := [...] int {1,2,3,5}
	//c := [...] int {1,2,3,4,5}
	d := [...] int {1,2,3,4}
	t.Log("a==b",a==b)
	t.Log("a==d",a==d)
	//t.Log("a==b",a==c)
}

func TestBitOperator(t *testing.T) {
	// &^ 按位置零运算符
	// 右边为1，则结果始终为零
	// 右边为0，则结果始终等于左边

	t.Log(0 &^ 0)
	t.Log(1 &^ 1)
	t.Log(1 &^ 0)
	t.Log(0 &^ 1)

	t.Log(4 &^ 5,5 &^ 4)
	t.Log(6 &^ 4,4 &^ 6)
}

const (
	READ = 1 << iota
	WRITE
	EXECUTE
)
func TestBitClear(t *testing.T) {
	var a int = 7 //0x0111
	t.Log(READ&a==READ,WRITE&a==WRITE,EXECUTE&a==EXECUTE)
	a = a &^ READ
	t.Log(READ&a==READ,WRITE&a==WRITE,EXECUTE&a==EXECUTE)

	a = a &^ EXECUTE
	t.Log(READ&a==READ,WRITE&a==WRITE,EXECUTE&a==EXECUTE)
}