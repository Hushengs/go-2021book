package main

import "fmt"

func main() {
	//源切片中元素类型为引用类型时，拷贝的是引用
	matA := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}
	matB := make([][]int, len(matA))
	copy(matB, matA)
	fmt.Printf("%p, %p\n", matA, matA[0]) // 0xc0000c0000, 0xc0000c2000
	fmt.Printf("%p, %p\n", matB, matB[0]) // 0xc0000c0050, 0xc0000c2000

	//正确拷贝多维数组
	matC := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}
	matD := make([][]int, len(matC))
	for i := range matC {
		matD[i] = make([]int, len(matC[i])) // 注意初始化长度
		copy(matD[i], matC[i])
	}
	fmt.Printf("%p, %p\n", matC, matC[0]) // 0xc00005c050, 0xc000018560
	fmt.Printf("%p, %p\n", matD, matD[0]) // 0xc00005c0a0, 0xc0000185c0
}
