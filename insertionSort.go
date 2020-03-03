package main 

import "fmt"

func insertionSort(arr[] int, n int){
	var i, key, j int
	for j = 1; i < n; i ++{
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key{
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main(){
	var array = [] int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Input Array = ", array)
	insertionSort(array, 11)
	fmt.Println("Sorted Array = ", array)
}