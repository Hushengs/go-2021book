package main

import "fmt"

func main() {
	// c := []int{1, 2, 3}
	// fmt.Println(c)
	// fmt.Println(len(c), cap(c))
	// d := c[:2]
	// fmt.Println(d)
	// fmt.Println(len(d), cap(d))
	// e := make([]int, 2, 3)
	// fmt.Println(e)
	// fmt.Println(len(e), cap(e))

	// //slice最近一个元素
	// a := []int{1, 2, 3}
	// fmt.Println(a)
	// a = append(a, 0)
	// lena := copy(a[1:], a[0:])
	// fmt.Println(a)
	// fmt.Println(lena)
	// a[0] = 4
	// fmt.Println(a)

	// //slice operate
	// opSlice := []int{1, 2, 3}
	// //删除末尾N个元素
	// N := 1
	// i := 1
	// a = a[:i+copy(a[i:], a[i+N:])]
	// fmt.Println(a)
	// opSlice = opSlice[:len(opSlice)-N]
	// //删除开头N个元素
	// opSlice = opSlice[N:]
	// opSlice = append(opSlice[:0], opSlice[N:]...)
	// opSlice = opSlice[:copy(opSlice, opSlice[N:])]
	// //删除中间N元素
	// opSlice = append(opSlice[:i], opSlice[i+N:]...)
	// opSlice = opSlice[:i+copy(opSlice[i:], opSlice[i+N:])]

	s := []int{4,5,6}
	fmt.Println(s[:0])
}
