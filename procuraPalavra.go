package main

import "fmt"

func main(){
	var i, j int
	var a string = "Isso é um exemplo de texto0"
	var p string = "exemplo0"

	for i=0; i<len(a)-1; i++{
		for j=0; j<len(p)-1; j++{
			if a[i+j] != p[j]{
				break
			}
		}
		if p[j] == 48{
			fmt.Printf("%d - %c\n", i, a[i])
		}
	}
}