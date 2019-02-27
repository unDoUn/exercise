package main

import "fmt"

/*
Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

==8ms
用map边循环边把（target-当前值）放入，之后若当前循环值存在直接可以找出
*/
func twoSum(nums []int, target int) []int {
	result:=[]int{0,0}
	contain:=make(map[int]int)
	for k,v:=range nums{
		if _,ok:=contain[v];ok{	//验证是否包含常用方法
			result[0]=contain[v]
			result[1]=k
			break
		}
		contain[target-v]=k
	}
	//fmt.Println(contain)
	return result
}


func main(){
	a:=[]int{1,3,4,8,9,3}
	b:=6
	fmt.Println(twoSum(a,b))
}
