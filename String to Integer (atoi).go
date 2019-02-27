package main

import (
	"fmt"
	"strconv"
)
/*String to Integer(atoi)

Implement 【atoi】 which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.*/


/*-------------第一次尝试
func myAtoi(str string) int {

	if len(strings.TrimSpace(str))==0 ||strings.TrimSpace(str)=="+"||strings.TrimSpace(str)=="-"{
		return 0
	}

	number:=[]string{"-","+","0","1","2","3","4","5","6","7","8","9"}
	//nonSpace:=strings.TrimSpace(str)
	flag:=false
	firstB:=false
	lastB:=false
	floatB:=false
	var res string

	var first,last int
	for k,v:=range str{
		if string(v)!=" " && !firstB{
			first=k
			firstB=true
		}
		if string(v)==" " && firstB && !lastB{
			last=k
			lastB=true
		}
		if string(v)=="."{
			floatB=true
		}
	}
	if !lastB{
		last=len(str)
	}
	fmt.Println("----",first,last)
	for _,v:=range number{  //判断开头是否为数字
		if str[first:first+1]==v{
			flag=true
			break
		}
	}
	if !flag{
		return 0
	}
	//var f bool
	res=str[first:last]
	/*if str[0:1]=="-"{
		f=true
	}
	var temp string
	if string(res[0])=="-" || string(res[0])=="+"{
		temp=string(res[0])
		res=res[1:]
	}else{
		temp=""
	}
	i:=0
	var ff bool
	for string(res[i])=="0" {
		i++
		if i==len(res){
			ff=true
			break
		}
	}
	if !ff{
		res=res[i:]
	}
	res=temp+res

	var result int
	if floatB{
		resultF,_:=strconv.ParseFloat(res,32)
		result=int(resultF)
	}else {
		resultI, _ := strconv.ParseInt(res, 10, 32)
		result=int(resultI)
	}
	//fmt.Println(res,"----",result)

	return result
}
*/
/*写了半天，按照出错实例修改越改越乱，循环越加越多，总体乱了思路也乱了*/

func main(){
	fmt.Println(myAtoi("     0-1   "))
}

func myAtoi(str string) int {
	var result int
	var temp string
	notStart:=false
	label:=false
	isFloat:=false
	temp=""
	trimeSpace:=false

	for _,v:=range str{

		if string(v)==" " && !notStart && !trimeSpace{
			continue
		}
		trimeSpace=true
		if string(v)!="-" && string(v)!="+"&&(string(v)<"0" || string(v)>"9")&& !notStart{
			return 0
		}
		if (string(v)=="-" || string(v)=="+") && !notStart && !label {
			temp+=string(v)
			label=true
			continue
		}
		/*if string(v)=="0" && !notStart {   //[1]
			continue
		}
		*/
		if (string(v)<"0" || string(v)>"9") && !notStart{
			return 0
		}

		notStart = true
		if string(v)<"0" || string(v)>"9"{
			break
		}
		temp+=string(v)
	}
	for _,v:=range temp{
		if string(v)=="."{
			isFloat=true
		}
	}

	if isFloat{
		resultF,_:=strconv.ParseFloat(temp,32)
		result=int(resultF)
	}else{
		resultI,_:=strconv.ParseInt(temp,10,32)
		result=int(resultI)
	}
	return result
}



/*写了半天，按照出错修改越改越乱，感觉要理下题目(依照循环顺序)：
1、前缀有空格都跳过,前缀或"-""+"后有0也跳过（"  -0012a42" 返回-12）
2、只能从"-"、"+" 还有数字开头，以其他开头均返回0
3、开头之后，内容只有数字，遇到字母就返回了
4、如果溢出，返回对应的最大最小值（-2^31 ～ 2^31-1）（这里strconv.ParseFloat/Int取int32遇到溢出会自动返回临界值）

-----开始是针对问题来改，有一个要求就改一个，用了几个循环，越来越乱；
！！!所以要注意，如果以一个对字符串的循环，会直观很多,然后乱的时候停下来理一遍思路再继续

出错：输入" +0 123 "   " -42   "  "   0-1 "
if顺序和是否能够进入的bool函数都很重要
而[1]其实不需要，如果前缀有零，strconv.ParseFloat/Int后也就没有了，否则加上的话，输入"   0-1 "会得到-1，正确的是0
---finally---试了很多次，错了很多次，才把循环里面的顺序弄清楚
4 ms
*/