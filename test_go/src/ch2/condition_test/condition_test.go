package condition_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMultiSec(t *testing.T) {
	if a:=1==1; a {
		t.Log("a==1")
	} else {
		t.Log("a!=1")
	}

	if a, b := div(10,2); b==""{
		t.Log(a)
	} else {
		t.Log(b)
	}
}

func div(i int, j int ) (int, string) {
	if j==0 {
		return -1,"divide by zero"
	} else {
		return  i/j, ""
	}
}

func TestSwitch(t *testing.T)  {
	switch os := runtime.GOOS;  os {
	case "1","darwin": //可以逗号分割多条件
		fmt.Println("MacOS")
	//	默认break
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n",os)
	}

	for i := 0; i<5; i++ {
		switch i {
		//相当于 if。。。else if ...else 语句
		case 0,2:
			fmt.Println("Even")
		case 1,3:
			fmt.Println("Odd")
		default:
			fmt.Println("not 0-3")
		}
	}


}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i<11; i++ {
		switch  {
		//相当于 if。。。else if ...else 语句
		case i>=0 && i<=3:
			fmt.Println("(0,3]")
		case i>3 && i<=5:
			fmt.Println("(3,5]")
		case i>5 && i<10:
			fmt.Println("(5,10]")
		default:
			fmt.Println(">10")
		}
	}
}