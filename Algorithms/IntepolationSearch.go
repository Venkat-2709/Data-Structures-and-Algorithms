package main

import "fmt"

func interpolationSearch(array []int, n int, value int) int {
	first := 0
	last := n - 1

	for first <= last && value >= array[first] && value <= array[last] {
		if first == last {
			if array[first] == value {
				return first
			}
			return -1
		}

		position := first + (int(last-first) / (array[last] - array[first]) * (value - array[first]))

		if array[position] == value {
			return position
		}

		if array[position] < value {
			first = position + 1
		} else {
			last = position - 1
		}
	}
	return -1
}

func main() {

	var array = []int{15, 48, 69, 78, 102, 119, 156, 189, 211, 268, 300, 342, 385, 401}
	value := 211

	position := interpolationSearch(array, len(array), value)

	if position == -1 {
		fmt.Printf("\nElement %d in not in the list", value)
	} else {
		fmt.Printf("\nElement %d is found at index %d\n", value, position)
	}
}
