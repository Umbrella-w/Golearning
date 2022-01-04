## Slice（切片）

### 切片的本质

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组`a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]`，相应示意图如下。

![slice_01](https://www.liwenzhou.com/images/Go/slice/slice_01.png)

切片`s2 := a[3:6]`，相应示意图如下：![slice_02](https://www.liwenzhou.com/images/Go/slice/slice_02.png)

