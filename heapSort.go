package main

import "fmt"

func heapSort(a [] int, n int){
	var i = n / 2
	var pai, filho, t int
	for true{
		if i > 0{
			i--
			t = a[i]
		} else {
			n--
			if n <= 0 {
				return
			}
			t = a[n]
			a[n] = a[0]
		}
		pai = i
		filho = i*2 + 1
		for filho < n{
			if filho + 1 < n && a[filho+1] > a[filho]{
				filho ++
			}
			if a[filho] > t {
				a[pai] = a[filho]
				pai = filho
				filho = pai * 2 + 1
			} else {
				break
			}
		}
		a[pai] = t
	}
}

func main(){
	var array = [] int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Input Array = ", array)
	heapSort(array, 11)
	fmt.Println("Input Array = ", array)
}