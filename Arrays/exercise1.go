/*
1、求数组[1, 3, 5, 7, 8]所有元素的和
2、找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
 */

package main

import "fmt"

func main() {
	//1:求数组[1, 3, 5, 7, 8]所有元素的和
	a := [...]int{1,3,5,7,8}
	fmt.Println(a)
	sum := 0
	for _,value := range a{
		sum += value	//等价于 sum = sum + value

	}
	fmt.Printf("sum=%d\n",sum)

	findValue()		//2、找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	findValue1()	//不同写法
	findValue2()
}

func findValue()  {
	a := [5]int{1, 3, 5, 7, 8}
	for i, j := range a {
		for l, k := range a[i+1:] {
			if j+k == 8 {
				fmt.Printf("(%d, %d)\n", i, l+i+1)
			}
		}
	}
}

func findValue1(){
	a := [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			if a[i] + a[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}

func findValue2() {
	var array = [5]int{1, 3, 5, 7, 8}
	length := len(array)
	const num = 8
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if array[i]+array[j] == num {
				fmt.Printf("(%d, %d) \n", i, j)
			}
		}
	}
}