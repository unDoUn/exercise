package main

import "fmt"
/*
5 Longest Palindromic Substring
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"*/

/*func longestPalindrome(s string) string {
	if len(s)==0{
		return ""
	}
	var max int
	var resultStr [2]int
	for k:=range s{
		var flag bool
		index:=strings.LastIndexAny(s[k+1:],s[k:k+1])		//[1、2]
		tail:=index+k+1
		for k+1 <= len(s)-1 && index != -1{    //取最后一个索引循环
				//fmt.Println(s[k+1:],s[k:k+1],index,"-----",tail)
				head1 := k
				tail1:=tail
				var count int
				for count=0; head1 < tail1; head1, tail1 = head1+1, tail1-1 { //判断是否回文循环
					if s[head1:head1+1] == s[tail1:tail1+1] {
						count++
					}
				}
				flag =2*count >= (tail - k)

				if flag && (tail-k) > max{
					max = tail - k
					resultStr[0] = k
					resultStr[1] = tail
				}
				index=strings.LastIndexAny(s[k+1:tail],s[k:k+1])
				tail=index+k+1
			}
		}

	return s[resultStr[0]:resultStr[1]+1]
}
*/


/*思想不复杂，但是关于有两个细节问题：
1、range循环若设为k,v，直接打印v，得到的并不是字符而是对应的ascii码，需要转换
（两种方法：切片取部分 s[k:k+1]或者string转换 string(rune(v))）
2、[2]处取包含的最后一个值的索引时用了切片 s[k+1:]，此时(k+1)处即为索引0对应的位置，
所以要加上之前的k个才能对应上原来的位置
3、特殊值处理：如"a"、"qwer"类没有相同值的，返回第一个值；""空值返回空

报错-超时 Time Limit Exceeded（三次循环了）
*/

var max int
var res [2]int
var maxStr string


func longestPalindrome(s string) string {

	if len(s)==0{
		return ""
	}
	res=[2]int{0}     //[3]
	max=0				//[4]
	for k:=range s{
		findLongestString(s,k,k)
		if k+1<len(s) {
			findLongestString(s, k, k+1)
		}
	}
	return  maxStr
}

func findLongestString(s string,h int,t int) {
	var count int
	var flag bool
	for h>=0 && t<len(s){
		if h!=t && s[h:h+1]==s[t:t+1]{
			count++
		}else if s[h:h+1]!=s[t:t+1]{
			break
		}
		h,t=h-1,t+1
	}
	h,t=h+1,t-1
	if t-h ==0{
		flag=false
	}else {
		flag = 2*count >= (t - h)
	}
	if t-h>max && flag{
		max=t-h
		res[0]=h
		res[1]=t
	}

	maxStr=s[res[0]:res[1]+1]
}

func main(){
	//s:="qwertyui"
	//s:="bb"
	//s:="ac"
	s:="cbbd"
	fmt.Println(longestPalindrome(s))

}

/*若字符串为"cbbd"，如果只看head,tail均为k，则只会返回奇数子串，这里就会认为不回文，返回c
所以也要令head,tail为k，k+1，返回偶数串情况
所以单独用一个函数拿出来弄
[3、4]runtime error---"a"  "ac"   "bb"等都会报错
查了下说一般没有初始化，运行时是随机数；以为返回是空指针，其实返回了一个野指针
[4]如果不要，运行 "cbbd"，出来结果是c，虽然golang上结果正确
最后  28 ms*/