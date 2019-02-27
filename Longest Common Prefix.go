package main

import "fmt"

/*
14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/


func longestCommonPrefix(strs []string) string {
	shortStr:=Shortest(strs)
	for i:=0;i<len(shortStr);i++{
		m:=shortStr[i]
		for _,v:=range strs{
			if m!=v[i]{
				return shortStr[:i]
			}
		}
	}
	return shortStr
}

func Shortest(s []string) string{
	var temp string
	if len(s)==0{
		return ""
	}
	temp=s[0]
	for _,v :=range s{
		if len(v)<len(temp){
			temp=v
		}
	}
	return temp
}






func main(){
	ss:=[]string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(ss))
}

/*
重点：字符串---类比一位数组/切片   字符串切片--类比二维数组/切片处理
注意：
1、s:="cfvfsv"  s[0]是字符c对应的ASCII码，string(s[0])才是字符c
2、len(s)==0表示切片为空，s==nil不表示切片为空

Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Longest Common Prefix.
*/