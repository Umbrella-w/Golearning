```
/*
Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：
var 数组变量名 [元素数量]type

[5]int和[10]int是不同的类型

数组可以通过下标进行访问，下标是从`0`开始，最后一个元素下标是：`len-1`，访问越界（下标在合法范围之外），则触发访问越界，会panic
 */
package main

import "fmt"

func main() {
   var a [3]int   //初始化值为[0 0 0]
   var b [4]int   //初始化值为[0 0 0 0]
   //a = b //不可以这样做，因为此时a和b是不同的类型
   fmt.Println(a,b)

//数组的初始化
//方法一：
   var testArray [3]int                        //数组会初始化为int类型的零值
   var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
   var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
   fmt.Println(testArray)                      //[0 0 0]
   fmt.Println(numArray)                       //[1 2 0]
   fmt.Println(cityArray)                      //[北京 上海 深圳]


//方法二：

   var numArray1 = [...]int{1, 2}
   var cityArray1 = [...]string{"北京", "上海", "深圳"}

   fmt.Println(numArray1)                           //[1 2]
   fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
   fmt.Println(cityArray1)                          //[北京 上海 深圳]
   fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string

//方法三：
   c := [...]int{1: 1, 3: 5}
   fmt.Println(c)                  // [0 1 0 5]
   fmt.Printf("type of a:%T\n", c) //type of a:[4]int

//数组的遍历
//遍历数组test有以下两种方法：
   var test = [...]string{"北京","上海","深圳"}
   //方法1：for循环遍历
   for i := 0; i < len(test);i++{
      fmt.Println(test[i])   //把下标以及对应值进行遍历输出
      //北京
      //上海
      //深圳

   }
   //方法2：for range 遍历
   for index ,value := range test {
      fmt.Println(index,value)   //把下标以及对应值进行遍历输出
      //0 北京
      //1 上海
      //2 深圳
   }


}
```