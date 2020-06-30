package CourseThree

import (
	"fmt"
	"time"
)

func Race() {
	number := 0
	/*
		Even though the number should be 0 after Inc and Decr but in the case of
		race condition they both tries to get it and and they both take 0. Even Inc take first
		it can not handle to increment it and decrement can not handle to decrement and
		in result it is unknown what the result will be
	*/
	go Inc(&number)
	go Decr(&number)
	time.Sleep(3000)
}

func Inc(number *int) {
	*number++
	fmt.Printf("Incremented! : %d\n", *number)
}

func Decr(number *int) {
	*number--
	fmt.Printf("Decremented! : %d\n", *number)
}
