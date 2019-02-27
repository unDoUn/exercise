package main

import (
	"fmt"
)

/*
20. Valid Parentheses

Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/


func isValid(s string) bool {
	if len(s)==0{
		return true
	}
	var zhan []string
	for _,v:=range s{
		if string(v)=="(" || string(v)=="{" || string(v)=="["{
			zhan=append(zhan,string(v))
		}else {
			isMatch:=Output(zhan,string(v))
			if isMatch {
				zhan=zhan[:len(zhan)-1]
			}else{
				return false
			}
		}
	}

	if len(zhan)==0{
		return true
	}else{
		return false
	}
}


func Output(zhan []string,now string) bool{
	n:=len(zhan)
	if n==0{
		return false
	}
	sym:=map[string]string{
		"}":"{",
		"]":"[",
		")":"(",
	}
	if sym[now]==zhan[n-1]{
		return true
	}else{
		return false
	}
}



func main(){
	s:="]"
	fmt.Println(isValid(s))
}


/*
字符串：s:="dfaf"  s[0]为d的ASCII码,而s[0:1]为d
易忽视：只要有对切片/数组等数据结构长度进行变化的时候，
都要提前考虑到长度可能超过少于正常范围
如：Output()函数中，zhan[n-1]前要进行长度是否为0的判断
if n==0{
		return false
	}
总结：这里利用切片来构成了一个栈的结构
其中入栈--append函数   出栈--切片自身的[:len()-1]

Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Valid Parentheses.
*/