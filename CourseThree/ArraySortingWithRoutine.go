package CourseThree

import (
	"fmt"
	"sync"
)

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.Each partition should be of approximately
equal size. Then the main goroutine should merge the 4 sorted sub arrays into one large sorted array.
The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of
the array should print the subarray that it will sort. When sorting is complete, the main goroutine
should print the entire sorted list.

*/

// Sorts
func Sort(elementList []int) {
	var wg sync.WaitGroup
	lenOfArray := len(elementList)
	index := 0
	PrintArray(elementList)
	if lenOfArray < 4 {
		wg.Add(1)
		go SortPartOfArray(&elementList, 0, lenOfArray, &wg)
	} else {
		p1Len, p2Len, p3Len, p4Len := LenDivideLenIntoParts(lenOfArray)
		wg.Add(4)
		go SortPartOfArray(&elementList, index, index+p1Len, &wg)
		index = index + p1Len
		go SortPartOfArray(&elementList, index, index+p2Len, &wg)
		index = index + p2Len
		go SortPartOfArray(&elementList, index, index+p3Len, &wg)
		index = index + p3Len
		go SortPartOfArray(&elementList, index, index+p4Len, &wg)
		wg.Wait()
		index = 0
		MergeTwoPart(&elementList, 0, elementList[index:index+p1Len], elementList[index+p1Len:index+p1Len+p2Len])
	}
	PrintArray(elementList)
}

func MergeTwoPart(elementList *[]int, index int, p1Part []int, p2Part []int) {
	for i, k, j := 0, 0, index; i < len(p1Part) && k < len(p2Part); j++ {
		if p1Part[i] <= p2Part[k] {
			fmt.Printf("%d<=%d\n", p1Part[i], p2Part[k])
			(*elementList)[j] = p1Part[i]
			i++
		} else {
			fmt.Printf("%d>=%d\n", p1Part[i], p2Part[k])
			(*elementList)[j] = p2Part[k]
			k++
		}
	}
}

//Sorts part of Array
func SortPartOfArray(partOfArray *[]int, firstIndex int, lastIndex int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	if len(*partOfArray) == 1 {
		PrintArray(*partOfArray)
		return *partOfArray
	}
	BubbleSort(partOfArray, firstIndex, lastIndex)
	PrintArray(*partOfArray)
	return *partOfArray
}

//Bubble sort implementations
func BubbleSort(slice *[]int, firstIndex int, lastIndex int) []int {
	hasSwapped := false
	for !hasSwapped {
		hasSwapped = true
		for i := firstIndex; i < lastIndex-1; i++ {
			if (*slice)[i] > (*slice)[i+1] {
				hasSwapped = false
				(*slice)[i], (*slice)[i+1] = (*slice)[i+1], (*slice)[i]
			}
		}
	}
	return *slice
}

// create "equal" parts length
func LenDivideLenIntoParts(lenOfArray int) (int, int, int, int) {
	lenOfParts := lenOfArray / 4
	partsLenList := []int{lenOfParts, lenOfParts, lenOfParts, lenOfParts}
	for k, i := 0, lenOfArray%4; i > 0; i, k = i-1, k+1 {
		partsLenList[k] += 1
	}
	return partsLenList[0], partsLenList[1], partsLenList[2], partsLenList[3]
}

func PrintArray(array []int) {
	fmt.Print("[")
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d,", array[i])
	}
	fmt.Printf("%d]\n", array[len(array)-1])
}
