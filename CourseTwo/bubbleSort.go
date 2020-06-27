package CourseTwo

import (
	"fmt"
)

func BubbleSort() {
	unsortedSlice := GetUserInput()
	sorted := Sort(unsortedSlice)

	fmt.Println("Unsorted slice :")
	PrintSlice(unsortedSlice)

	fmt.Println("Sorted slice :")
	PrintSlice(sorted)
}

func GetUserInput() []int {
	var unsortedSlice []int
	var userInput int
	// up to 10 element
	for i := 1; i <= 10; i++ {
		fmt.Print("Enter new integer :")
		fmt.Scan(&userInput)
		unsortedSlice = append(unsortedSlice, userInput)
		fmt.Println()
	}
	return unsortedSlice
}

func Sort(slice []int) []int {
	hasSwapped := false
	for !hasSwapped {
		hasSwapped = true
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				hasSwapped = false
				Swap(&slice[i], &slice[i+1])
			}
		}
	}
	return slice
}

func Swap(element *int, element2 *int) {
	temp := *element
	*element = *element2
	*element2 = temp
}

func PrintSlice(sorted []int) {
	for _, v := range sorted {
		fmt.Printf("[%d] \n", v)
	}
}
