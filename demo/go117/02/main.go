package main

/*
go run -gcflags=-G=3 02/main.go
*/

import (
	"fmt"
)

type vector[T any] []T

func printSlice[T any](s []T){
	for _,v := range s {
		fmt.Printf("%v ",v)
	}
	fmt.Print("\n")
}

func main(){
	v := vector[int]{58,188}
	printSlice(v)
	v2 := vector[string]{"红烧肉","九转大肠"}
	printSlice(v2)
	v3 := []float64{1.1,2.2}
	printSlice(v3)
}