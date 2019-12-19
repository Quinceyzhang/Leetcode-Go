package main

import "fmt"

//设计算法:
//运用map,统计nums1中值出现的次数-map[值]次数
//遍历nums2中的值,查看值是否在map中的出现

func intersectPlate(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var arr []int
	for _, v := range nums1 {
		m[v]++
	}
	fmt.Println(m)

	for _, v := range nums2 {
		times, ok := m[v] //v是nums2中的值,m[v]是map中的值.m[v]==times
		fmt.Printf("v=%d,times=%d\n", v, times)
		if ok && times > 0 {
			arr = append(arr, v)
			m[v]-- //所有出现的数字都+1,最后要减掉1
		}
	}
	return arr
}

func main() {
	a1 := []int{1, 2, 2, 2, 3, 4, 5, 6, 7}
	a2 := []int{2, 1, 3, 8}
	fmt.Println(intersect(a1, a2))
}
