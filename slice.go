package main

import (
	"fmt"
	"sort"
	"strconv"
)

func slice() {
	var userEntry string
	orderedSlice := make([]int, 1)
	for {
		fmt.Print("Enter the number to add slice : ")
		fmt.Scan(&userEntry)
		if userEntry != "X" {
			value, error := strconv.Atoi(userEntry)
			if error == nil {
				orderedSlice = append(orderedSlice, value)
				sort.Ints(orderedSlice)
			}
		} else {
			break
		}
		fmt.Println(orderedSlice)
	}
	fmt.Scan()

}
