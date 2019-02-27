package main

import "fmt"

/*ZigZag Conversion
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I


*/

func convert(s string, numRows int) string {
	n:=len(s)
	if s=="" || n==0 || numRows==0{
		return ""
	}
	if numRows==1{   //[1]
		return s
	}
	var res string
	size:=2*(numRows-1)
	for i:=0;i<numRows;i++{

		for j:=i;j<n;j+=size {
			res += s[j : j+1]
			if i != 0 && i != numRows-1 {
				slash := j + size - 2*i
				if slash<n {
					res += s[slash : slash+1]
				}
			}
		}
	}
	return res
}


func main(){
	fmt.Println(convert("A",4))

}

/*
就是画出字符位置找规律，找出规律后，代码很简单
 0 2 4 6

 1 3 5 7
-----------------
 0     4     8

 1  3  5  7  9

 2     6    10
----------------------
 0      6        12

 1   5  7    11  13

 2 4    8 10     14

 3      9        15

开始找出了规律，但是不知道怎么一边处理斜的部分，一边处理竖的（间隔不一样）
后来看博客，才发现可以吧斜着的和竖着的分开处理，因为竖着的每次间隔相等，斜着的每次间隔相等
i-当前行数  j-当前竖着的某字符位置
slash := j + size - 2*i 表示斜着字符所在的位置

报错 Time Limit Exceeded---（中止参数 "AB",1）---运行发现无限循环
加入[1]进行限制
16 ms
*/
