package main

import (
	"fmt"
	"math"
	"strconv"
)


/*
9. Palindrome Number
Determine whether an integer is a palindrome.
An integer is a palindrome when it reads the same backward as forward.
Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121.
From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left.
Therefore it is not a palindrome.
*/


func isPalindrome(x int) bool {
	if x<0{
		return false
	}
	n:=len(strconv.Itoa(x))
	var xSlice []int
	for i:=0;i<n;i++{
		m:=math.Mod(float64(x),10)
		x=x/10
		xSlice=append(xSlice,int(m))
	}
	for i:=0;i<n/2;i++{
		if xSlice[i] != xSlice[n-i-1]{
			return false
		}
	}
	return true
}


func main(){
	x:=12321
	fmt.Println(isPalindrome(x))

}

/*
60 ms
key:多用万能的切片
*/