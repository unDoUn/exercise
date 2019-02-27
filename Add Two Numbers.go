package main

import (
	"fmt"
	"math"
)

/*
2. Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3:=new(ListNode)
	head:=l3
	carry:=[]int{0}
	for l1!=nil && l2!=nil{
		node:=new(ListNode)
		add:=l1.Val+l2.Val+carry[len(carry)-1]
		node.Val=int(math.Mod(float64(add),10))
		fmt.Println("----",node.Val)
		l3.Next=node
		l3=node

		carry=append(carry,add/10)
		l1=l1.Next
		l2=l2.Next
	}

	for l1!=nil{
		fmt.Println("-----",l1.Val)
		l3.Next=l1
		l3=l1
		add:=l1.Val+carry[len(carry)-1]
		l1.Val=int(math.Mod(float64(add),10))
		carry=append(carry,add/10)
		l1=l1.Next
		if l1!=nil{
			l3=l3.Next
		}
	}
	if l2!=nil{
		l3.Next=l2
		l3=l2
		for  {
			add := l2.Val + carry[len(carry)-1]
			l2.Val = int(math.Mod(float64(add), 10))
			carry = append(carry, add/10)
			l2 = l2.Next
			if l2 != nil {
				l3 = l3.Next
			}else {
				break
			}
		}
	}
	if carry[len(carry)-1]!=0{
		nodeLast:=new(ListNode)
		nodeLast.Val=carry[len(carry)-1]
		nodeLast.Next=nil
		l3.Next=nodeLast
	}else{
		l3.Next=nil
	}
	return head.Next
}


func listLen(l *ListNode) int{
	var i int
	for l!=nil{
		i++
		l=l.Next
	}
	return i
}


func main(){
	l1:=&ListNode{9,&ListNode{1,&ListNode{Val:6}}}
	l2:=&ListNode{Val:0}
	l3:=addTwoNumbers(l1,l2)
	var s []int
	for l3!=nil{
		s=append(s,l3.Val)
		l3=l3.Next
	}
	fmt.Println(s,listLen(l1),listLen(l2))
}

/*
易忽视：
1、两个链表长短不一样的时候，长的那个列表依然要加进位项，
而且要一直循环加，因为多的部分可能是9999，即一直会进位
2、加入节点时，容易遗漏其中某一步：指针要连接到节点，并且自身也还要移动
   l3.Next=node
   l3=node

总结：
分成三个步骤形成新的链表
1 l1和l2值相加得到新链表
2 l1或者l2的剩余部分，l1/l2加上进位后接在后面
3 l1和l2都遍历完之后，如果不是0，还要加上进位的数接在后面

*/
