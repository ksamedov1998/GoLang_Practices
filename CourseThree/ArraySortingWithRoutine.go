package CourseThree

import (
	"fmt"
	"sync"
)

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
		MergeTwoPart(&elementList, index, CreateArray(elementList, index, index+p1Len), CreateArray(elementList, index+p1Len, index+p1Len+p2Len))
		index += p1Len + p2Len
		MergeTwoPart(&elementList, index, CreateArray(elementList, index, index+p3Len), CreateArray(elementList, index+p3Len, index+p3Len+p4Len))
		MergeTwoPart(&elementList, 0, CreateArray(elementList, 0, len(elementList)/2+len(elementList)%2), CreateArray(elementList, len(elementList)/2+len(elementList)%2, len(elementList)))
	}

	PrintArray(elementList)
}

func CreateArray(list []int, index int, lastIndex int) []int {
	var array []int
	for i, k := index, 0; i < lastIndex; i, k = i+1, k+1 {
		array = append(array, list[i])
	}
	return array
}

func MergeTwoPart(elementList *[]int, index int, p1Part []int, p2Part []int) {
	for i, k, j := 0, 0, index; i < len(p1Part) && k < len(p2Part); {
		var p1Changed bool
		if p1Part[i] <= p2Part[k] {
			p1Changed = true
			fmt.Printf("%d<=%d\n", p1Part[i], p2Part[k])
			(*elementList)[j] = p1Part[i]
			i++
		} else {
			p1Changed = false
			fmt.Printf("%d>=%d\n", p1Part[i], p2Part[k])
			(*elementList)[j] = p2Part[k]
			k++
		}
		j++
		if !(i < len(p1Part)) {
			if p1Changed {
				k--
			}
			for l := k; l < len(p2Part); l++ {
				(*elementList)[j] = p2Part[l]
				j++
			}
		} else if !(k < len(p2Part)) {
			if p1Changed {
				i--
			}
			for l := i; l < len(p1Part); l++ {
				(*elementList)[j] = p1Part[l]
				j++
			}
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
