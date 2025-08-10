package main

import (
	"errors"
	"fmt"
	"sort"
)

func secondMax(slice []int) (int, error) {
	if len(slice) < 2 {
		return 0, errors.New("there are less that two values in the slice")
	}

	var firstMax int = slice[0]
	var secondMax int = slice[1]

	if slice[0] < slice[1] {
		firstMax = slice[1]
		secondMax = slice[0]
	}

	for _, num := range slice[2:] {
		if num > firstMax {
			secondMax = firstMax
			firstMax = num
			continue
		}

		if num > secondMax && num != firstMax || firstMax == secondMax {
			secondMax = num
		}
	}

	return secondMax, nil
}

// func main() {
// 	fmt.Println(secondMax([]int{5, 7, 9, 11, 14, 16}))                                //14
// 	fmt.Println(secondMax([]int{16, 14, 11, 9, 7, 5}))                                //14
// 	fmt.Println(secondMax([]int{10, 20}))                                             //10
// 	fmt.Println(secondMax([]int{20, 10}))                                             //10
// 	fmt.Println(secondMax([]int{5, 5, 1}))                                            //1
// 	fmt.Println(secondMax([]int{10, 5, 10}))                                          //5
// 	fmt.Println(secondMax([]int{-10, -5, 0}))                                         //-5
// 	fmt.Println(secondMax([]int{1, 500, 2, 300, 3, 1000}))                            //500
// 	fmt.Println(secondMax([]int{-500, -50, -10, -2}))                                 //-10
// 	fmt.Println(secondMax([]int{-102308042342, 5349785493, 10, -1928439, 12948, 12})) //12948
// 	fmt.Println(secondMax([]int{5, 5, 5}))                                            //5
// }

func mergeAndSort(arr1, arr2 []int) []int {
	for _, arr := range arr2 {
		arr1 = append(arr1, arr)
	}

	sort.Ints(arr1)

	return arr1
}

// func main() {
// 	fmt.Println(mergeAndSort([]int{3, 1, 5}, []int{4, 2, 6})) // [1 2 3 4 5 6]
// 	fmt.Println(mergeAndSort([]int{8, 2, 0}, []int{7, 3, 1})) // [0 1 2 3 7 8]
// }

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println(isLeapYear(2020)) // true
// 	fmt.Println(isLeapYear(1900)) // false
// 	fmt.Println(isLeapYear(2000)) // true
// }

func maxOfThree(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else if c > b && c > a {
		return c
	}
	return a
}

// func main() {
// 	fmt.Println(maxOfThree(3, 7, 5))     // 7
// 	fmt.Println(maxOfThree(10, 2, 8))    // 10
// 	fmt.Println(maxOfThree(-1, -5, -10)) // -1
// 	fmt.Println(maxOfThree(10, 10, 10))  // 10
// 	fmt.Println(maxOfThree(10, 10, 11))  // 11
// 	fmt.Println(maxOfThree(10, 11, 10))  // 11
// 	fmt.Println(maxOfThree(11, 10, 10))  // 11
// 	fmt.Println(maxOfThree(1, 2, 80))    // 80
// }

func reverse(line string) string {
	length := len(line)

	intermediate := make([]int, length+1)

	for i, r := range line {
		intermediate[length-i] = int(r)
	}

	runes := make([]rune, length+1)
	for i, code := range intermediate {
		runes[length - i] = rune(code)
	}

	return string(runes)
}

func main() {
	fmt.Println(reverse("hello"))  // "olleh"
	fmt.Println(reverse("Привет")) // "тевирП"
}
