package main

import "fmt"

func main() {
//声明切片类型
	var a []string	//声明一个字符串切片
	var b = []int{}	//声明一个整型切片并初始化
	var c = []bool{false, true}	//声明一个布尔型切片并初始化
	var d = []bool{false, true}	//声明一个布尔型切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(d)
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较


	/*
	切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

	切片表达式
	切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。

	简单切片表达式
	切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的low和high表示一个索引范围（左包含，右不包含），也就是下面代码中从数组test中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	 */

	test := [5]int{1,2,3,4,5}
	s := test[1:3]	// s:= test[low:high],这里[1：3]取得是从数组下标为1开始到下标为2的值
	fmt.Printf("s:%v len(s):%v cap(s):%v\n",s,len(s),cap(s))	//s:[2 3] len(s):2 cap(s):4

	/*
	可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:

	a[2:]  // 等同于 a[2:len(a)]
	a[:3]  // 等同于 a[0:3]
	a[:]   // 等同于 a[0:len(a)]
	 */

//注意：
/*
   对于数组或字符串，如果0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。

   对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。常量索引必须是非负的，并且可以用int类型的值表示;对于数组或常量字符串，常量索引也必须在有效范围内。如果low和high两个指标都是常数，它们必须满足low <= high。如果索引在运行时超出范围，就会发生运行时panic。


*/

	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))	//s2:[5] len(s2):1 cap(s2):1





	/*
	完整切片表达式
	对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：

		a[low : high : max]
		上面的代码会构造与简单切片表达式a[low: high]相同类型、相同长度和元素的切片。另外，它会将得到的结果切片的容量设置为max-low。
		在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。
		完整切片表达式需要满足的条件是0 <= low <= high <= max <= cap(a)，其他条件和简单切片表达式相同。

	*/


		test1 := [5]int{1, 2, 3, 4, 5}
		t := test1[1:3:5]
		fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))	//t:[2 3] len(t):2 cap(t):4



//使用make()函数构造切片
	/*
	上面都是基于数组来创建的切片，如果需要动态的创建一个切片,需要使用内置的make()函数，格式如下：

	make([]T, size, cap)
	其中：
	T:切片的元素类型
	size:切片中元素的数量
	cap:切片的容量

	*/


	makeSlice := make([]int,2,10)
	fmt.Println(a)	//[]
	fmt.Println(len(makeSlice))	//2
	fmt.Println(cap(makeSlice))	//10

//判断切片是否为空


	/*
	切片不能直接比较，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
	切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	但是我们不能说一个长度和容量都是0的切片一定是nil
	 */


	/*
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	 */

//检查切片是否为空，始终使用len(s) == 0来判断，而不应该使用s == nil来判断。


//切片的赋值拷贝
	//下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	sliceA := make([]int,3)	//[0 0 0]
	sliceB := sliceA	//将sliceA直接赋值给sliceB,sliceA和sliceB共用一个底层数组
	sliceB[0] = 100
	fmt.Println(sliceA)	//[100 0 0]
	fmt.Println(sliceB)	//[100 0 0]

//切片遍历
	//切片的遍历方式和数组是一致的，支持索引遍历和for range遍历
	sTest := []int{1,3,5}
	for i:=0; i< len(sTest);i++{
		fmt.Println(i,sTest[i])
	}
	//等价于
	for index,value := range sTest{
		fmt.Println(index,value)
	}


//append()方法为切片添加元素
	/*
	Go语言的内建函数append()可以为切片动态添加元素。
	可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）
	 */

	var sTest1 []int
	sTest1 = append(sTest1,1)	//[1]
	sTest1 = append(sTest1,2,3,4)	//[1,2,3,4]
	sTest2 := []int{5,6,7,8}
	sTest1 = append(sTest1,sTest2...)	//[1 2 3 4 5 6 7 8]
	fmt.Println(sTest1)
//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化
	/*
	没有必要像下面的代码一样初始化一个切片再传入append()函数使用，

	s := []int{}  // 没有必要初始化
	s = append(s, 1, 2, 3)

	var s = make([]int)  // 没有必要初始化
	s = append(s, 1, 2, 3)
	 */

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice,i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	/*
	[0]  len:1  cap:1  ptr:0xc00000a180
	[0 1]  len:2  cap:2  ptr:0xc00000a190
	[0 1 2]  len:3  cap:4  ptr:0xc00000e1c0
	[0 1 2 3]  len:4  cap:4  ptr:0xc00000e1c0
	[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000102c0
	[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000102c0
	[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000102c0
	[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000102c0
	[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc00001a100
	[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc00001a100
	 */
	}



//使用copy()函数复制切片
	/*
	首先我们来看一个问题：

		func main() {
			a := []int{1, 2, 3, 4, 5}
			b := a
			fmt.Println(a) //[1 2 3 4 5]
			fmt.Println(b) //[1 2 3 4 5]
			b[0] = 1000
			fmt.Println(a) //[1000 2 3 4 5]
			fmt.Println(b) //[1000 2 3 4 5]
		}
		由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

		Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：

		copy(destSlice, srcSlice []T)
		其中：
		srcSlice: 数据来源切片
		destSlice: 目标切片
	 */


	srcSlice := []int{1,2,3,4,5}
	destSlice := make([]int,5,5)
	copy(destSlice, srcSlice )
	fmt.Println(srcSlice)	//[1 2 3 4 5]
	fmt.Println(destSlice)	//[1 2 3 4 5]
	destSlice [0] = 100
	fmt.Println(srcSlice)	//[1 2 3 4 5]
	fmt.Println(destSlice)	//[100 2 3 4 5]




//从切片中删除元素,Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。


	// 从切片中删除元素
	delteSv := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	delteSv = append(delteSv[:2], delteSv[3:]...)
	fmt.Println(delteSv) //[30 31 33 34 35 36 37]
//总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)





}
