package main

/*
go run -gcflags=-G=3 04/main.go
 */

import(
	"fmt"
	"strconv"
)

type Price int

type ShowPrice interface {
	String() string
}

func(i Price) String() string{
	return strconv.Itoa(int(i))
}

func showPriceList[T ShowPrice](s []T) (ret []string) {
	for index,v := range s {
		ret = append(ret, "第" + strconv.Itoa(index+1) + "菜的价格是：" + v.String())
	}
	return ret
}

func main(){
	fmt.Println(showPriceList([]Price{48, 88, 152, 219, 328}))
}
