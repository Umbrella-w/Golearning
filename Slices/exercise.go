package main

import (
	"fmt"
	"sort"
)

//1.写出下列代码的结果
func resultS() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("len(a):%d,cap(a):%d,value:%v\n",len(a),cap(a),a)	//[     0 1 2 3 4 5 6 7 8 9]
}


//2.使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序

func Sort() {
	var a = [...]int{3, 7, 8, 9, 1}
	// fmt.Println(a[:]) //将数组切片
	//sort包内部实现了内部数据类型的排序
	//升序排列
	sort.Ints(a[:])
	fmt.Println(a)
	//降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)
}

func main() {
	resultS()
	Sort()
}