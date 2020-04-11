package main

import (
	"fmt"
	"os"
)


func main() {
	fmt.Print("hello world "+f("John"))
	if len(os.Args) >1 {
		fmt.Println(": "+os.Args[1])
	}
	//os.Exit(0)
	fmt.Println(9/4)
}

func f(name string) string {
	return name
}