package main

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出: [4,9]
//说明：
//
//输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//我们可以不考虑输出结果的顺序。

//进阶:
//
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	var nums1 = []int{4, 9, 5}
	var nums2 = []int{9, 4, 9, 8, 4}
	println(intersect(nums1,nums2))
}
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var arr []int
	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		times, ok := m[v] //v是nums2中的值,m[v]是map中的值.m[v]==times
		if ok && times > 0 {
			arr = append(arr, v)
			m[v]-- //所有出现的数字都+1,最后要减掉1
		}
	}
	return arr
}
