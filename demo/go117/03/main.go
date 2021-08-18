package main

/*
go run -gcflags=-G=3 03/main.go
去掉列表中的string，再次编译运行，看看运行后是什么结果？
03\main.go:19:17: string does not satisfy Addable (string not found in int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, comple
x128)
*/

import "fmt"

type Addable interface{
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32,
	uint64, uintptr, float32, float64, complex64, complex128, string
}
func add[T Addable](a,b T) T {
	return a + b
}

func main(){
	fmt.Println(add(3,4))
	fmt.Println(add("Go","lang"))
}
