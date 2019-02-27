/*Median of two sorted arrays
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5*/


package main
import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m:=len(nums1)
	n:=len(nums2)
	var isOdd bool
	if (n+m)%2==1{
		isOdd=true
	}
	med:=(m+n)/2

	sum:=make([]int,0,m+n)
	var i,j int
	for i<m && j<n{
		if nums1[i]>nums2[j]{
			sum=append(sum, nums2[j])
			j++
		}else if nums1[i]<nums2[j]{
			sum=append(sum,nums1[i])
			i++
		}else{
			sum=append(sum,nums1[i],nums2[j])
			i++
			j++
		}
	}
	for i<m{
		sum=append(sum,nums1[i])
		i++
	}
	for j<n{
		sum=append(sum,nums2[j])
		j++
	}
	fmt.Println("------ ",sum,med)
	if isOdd{
		return float64(sum[med])
	}else{
		return float64(sum[med-1]+sum[med])/2
	}
}


func main(){
	a:=[]int{1,2}
	b:=[]int{3}
	fmt.Println(findMedianSortedArrays(a,b))
}


/* 36 ms
由于两个数组已经排好序，将其合并排序（每次取首位比较）取中位数即可
思路简单，但要对append()函数用法熟悉*/