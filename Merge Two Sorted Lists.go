package main

import (
	"fmt"
)

/*
21.Merge Two Sorted Lists
Merge two sorted linked lists and return it as a new list.
	The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode1 struct{
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	l3:=new(ListNode)
	head:=l3
	for l1!=nil && l2!=nil{
		node:=new(ListNode)
		if l1.Val<=l2.Val {
			node.Val=l1.Val
			l3.Next=node
			l3=node
			l1=l1.Next
		}else {
			node.Val=l2.Val
			l3.Next=node
			l3=node
			l2=l2.Next
		}
	}

	if l1!=nil{
		l3.Next=l1
	}else if l2!=nil{
		l3.Next=l2
	}else{
		l3.Next=nil
	}
	return head.Next
}



func main(){
	ss:=&ListNode{2,nil}
	l1:=new(ListNode)
	l1.Val=1
	l1.Next=&ListNode{2,&ListNode{Val:4}}
	l2:=new(ListNode)
	l2.Val=1
	l2.Next=&ListNode{3,&ListNode{Val:4}}
	var s []int
	l3:=mergeTwoLists(l1,l2)
	for l3!=nil {
		s=append(s,l3.Val)
		l3=l3.Next
	}
	fmt.Println(s,ss)

}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.6 MB, less than 45.54% of Go online submissions for Merge Two Sorted Lists.

总结：链表没怎么用过，耗时较长
链表特性：
1、链表要注意头节点的使用，因为最开始的指针如果一直向后移就不能代表那个链表了
2、对链表结构，最先应该想到的循环方式应该是检验链表下一位是否为空，
而不是根据长度进行循环，因为随着链表逐渐移动，长度也会相应的变化
3、结构体的声明和初始化格式
type student struct{
	name string
	age int
}

	1)用new来初始化
	s:=new(student)
	s.name="atgag"
	s.age=23

	2)用&初始化--和用new()等价
	s:=&student{"fargar",34}   //这样赋值要按顺序写，而且全都要赋值
或者 s:=&student{age:34}    //这样可以不按顺序，也可以只赋值某一部分

	3)用var声明结构体
	var ss student
	ss=student{"dvgrag",23}

	4)直接声明初始化也可以，但用的不多，一般还是加初始化
	ss:=student{"dvgrag",23}

*/