package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
7 Reverse Integer
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
*/

func reverse(x int) int {
	var xSlice []int
	var y,end,k int
	k =1
	n:=len(strconv.Itoa(x))

	if x<0 && n>0{
		end=n-1
	}else{
		end=n
	}

	for i:=0;i<end;i++ {
		m := math.Mod(float64(x), 10)
		x = x / 10
		xSlice = append(xSlice, int(m))
	}
	for i:=end-1;i>=0;i--{
		y+=xSlice[i]*k
		k=k*10
	}
	if y< -int(math.Pow(2,31)) || y>int(math.Pow(2,31))-1{
		return 0
	}
	return y
}


func main(){
	x:=1534236469
	fmt.Println(reverse(x))

}

/*
4 ms
1、易错：当x<0时，需要改结尾end的长度(n-1),否则负号会占一个位置影响结果
2、次方写法：math.Pow(x,n)   取余数：math.Mod(x,y) 且x,y要是float64类型
3、strconv.Itoa(x) 整型int变成字符串类型，可以通过转换成字符串得到整型长度
4、切片自带 append() 函数
*/