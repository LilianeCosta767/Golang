package main

import "fmt"

func quickSort(values [] int, began int, end int){
	var i, j, pivo, aux int
	i  = began
	j = end - 1
	pivo = values[(began+end)/2]
	for i <= j{
		for values[i] < pivo && i < end{
			i++
		}
		for values[j] > pivo && j > began{
			j--
		}
		if i <= j{
			aux = values[i]
			values[i] = values[j]
			values[j] = aux
			i++
			j--
		}
	}
	if j > began{
		quickSort(values, began, j+1)
	}
	if i < end{
		quickSort(values, i, end)
	}
}

func main(){
	var array = [] int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Input Array = ", array)
	quickSort(array, 0, 11)
	fmt.Println("Input Array = ", array)
}