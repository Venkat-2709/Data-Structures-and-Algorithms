package main

import "fmt"

func getMaxValue(arr []int) int {
	maxValue := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	return maxValue
}

func sort(arr []int, n int, exponent int) {
	var outputArr [10000]int
	var countArr [10]int

	for i := 0; i < n; i++ {
		countArr[(arr[i]/exponent)%10] += 1
	}

	for i := 1; i < 10; i++ {
		countArr[i] += countArr[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		outputArr[countArr[(arr[i]/exponent)%10]-1] = arr[i]
		countArr[(arr[i]/exponent)%10] -= 1
	}

	for i := 0; i < n; i++ {
		arr[i] = outputArr[i]
	}
}

func radixSort(arr []int, n int) {
	maxValue := getMaxValue(arr)

	for i := 1; maxValue/i > 0; i *= 10 {
		sort(arr, n, i)
	}
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d \t", arr[i])
	}
	fmt.Printf("\n")
}

func main() {
	var arr = []int{456, 354, 23, 586, 12, 2360, 369, 485, 75, 985, 22}

	printArr(arr)
	radixSort(arr, len(arr))
	printArr(arr)
}
