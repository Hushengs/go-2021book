package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//1.3.1
	s := "hello, world"
	// hello := s[:5]
	// world := s[7:]
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len)

	//byte原字节嘛
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	//遍历UTF8解码后的Unicode码点值
	for i, c := range "世界abc" {
		fmt.Println(i, c)
	}

	//int32
	for i, c := range []rune("世界abc") {
		fmt.Println(i, c)
	}

	//int
	for i, c := range []rune("世界abc") {
		fmt.Println(i, string(c))
	}

}
