package main

import (
	"fmt"
)

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		d := 0
		_, _ = fmt.Scanf("%d", &d)
		arr[i] = d
	}
	//quickSort(arr)
	bubbleSort(arr)
	printAll(arr)
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	n := len(arr)

	for true {
		p := arr[n-1]
		i, j := 0, n-1
		for ; i < n - 1 && arr[i] <= p; i++ {
		}
		for ; j > 0 && arr[j] > p; i-- {
		}
		if i >= j {
			arr[i], arr[n-1] = arr[n-1], arr[i]
			quickSort(arr[:i])
			quickSort(arr[i+1:])
			break
		} else {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for  j := 0; j < n - i - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func printAll(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	println()
}
