package main

import "fmt"

func main() {
	var a = []interface{}{123, "abc"}
	//传入解包a
	Print(a...)
	//传入非解包a
	Print(a)

	//defer延迟，闭包的这种以引用方式访问外部变量的行为可能会导致一些隐含的问题,闭包对捕获的外部变量并不是以传值方式访问，而是以引用方式访问3,3,3
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}

	//正确处理 2,1,0
	for i := 0; i < 3; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}
