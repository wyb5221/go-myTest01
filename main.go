package main

import (
	"fmt"
	task1 "go-myTest01/task"
)

func main() {

	i := 5
	task1.Test1(i)
	fmt.Println("调用test1后的值i：", i)
	task1.Test2(&i)
	fmt.Println("调用test2后的值i：", i)
}
