package main

import "fmt"

func trunc() {
	var variable float32
	fmt.Println("Enter the floatin number :")
	fmt.Scan(&variable)
	fmt.Println(int(variable))
	fmt.Scan()
}
