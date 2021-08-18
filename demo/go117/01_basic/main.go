package main

import "fmt"

/*
go run -gcflags=-G=3 01_basic/main.go
 */
func printSlice[T any](s []T){
	for _,v := range s {
		fmt.Printf("%v ",v)
	}
	fmt.Print("\n")
}

func main(){
	printSlice[int]([]int{66,77,88,99,100})
	printSlice[float64]([]float64{1.1,2.2,3.3,4.4,5.5})
	printSlice[string]([]string{"红烧肉","清蒸鱼","大闸蟹","九转大肠","重烧海参"})
	printSlice([]int64{55,44,33,22,11})
}
