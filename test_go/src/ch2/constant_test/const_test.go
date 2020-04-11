package constant_test

import (
	"testing"
)
const WEEK = 7
const (
	Mon = 0 + iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

const (
	READ = 1 << iota
	WRITE
	EXECUTE
)

func TestConstantWeek(t *testing.T) {
	t.Log(Mon,Tue,Sat,Sun)
}

func TestConstantStatus(t *testing.T) {
	t.Log(READ,WRITE,EXECUTE)
	var a int = 7 //0x0111
	t.Log(READ&a==READ,WRITE&a==WRITE,EXECUTE&a==EXECUTE)
	a = 6 //0x0110
	t.Log(READ&a==READ,WRITE&a==WRITE,EXECUTE&a==EXECUTE)
}