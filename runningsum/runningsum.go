package main

import "fmt"

func main() {
	var nums = []int{1, 1, 1, 1, 1}
	var result []int = runningSum(nums)
	var optResult []int = runningSumOptimize(nums)
	output(result, "Normal Algo")
	output(optResult, "Optimized Algo")
}

func output(result []int, desc string) {
	fmt.Print(desc, "\n")
	for i := 0; i < len(result); i++ {
		if result[i] != len(result) {
			fmt.Print(result[i], ",")
		} else {
			fmt.Print(result[i], "\n")
		}
	}
}

func runningSum(arr []int) []int {
	var newArr []int = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			newArr[j] = newArr[j] + arr[i]
		}
	}
	return newArr
}

func runningSumOptimize(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	return arr
}
