package main

//给定一个长度为 n 的整数数组，你的任务是判断在最多改变 1 个元素的情况下，该数组能否变成一个非递减数列。
//
//我们是这样定义一个非递减数列的： 对于数组中所有的 i (1 <= i < n)，满足 array[i] <= array[i + 1]。
//
//示例 1:
//
//输入: [4,2,3]
//输出: True
//解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
//示例 2:
//
//输入: [4,2,1]
//输出: False
//解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
//说明:  n 的范围为 [1, 10,000]。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/non-decreasing-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//var nums = []int{1, 2, 6, 7, 8, 8, 7}
	//var nums = []int{1, 1, 2, 3}
	//var nums = []int{4, 2, 3}
	//var nums = []int{1, 5, 4, 6, 7, 10, 8, 9}
	//println(largerNumber(6, 8, nums))
	var nums = []int{3, 3, 2, 2}
	//var nums = []int{1, 3, 2, 3, 3}
	//items := largerNumber(6, 8, nums)
	//for _, v := range items {
	//	println(v)
	//}
	println(checkPossibility(nums))
}

//func checkPossibility(nums []int) bool {
//	var length = len(nums)
//	if length == 1 {
//		return true
//	}
//	var items []int
//	var gMap = make(map[int]int)
//
//	var gMap1 = make(map[int]int)
//	for _, v := range nums {
//		gMap1[v]++
//	}
//	for index, value := range nums {
//		items = append(items, largerNumber(index, value, nums)...)
//		for _, v := range items {
//			gMap[v]++
//		}
//	}
//	if len(gMap) > 1 {
//		return false
//	} else {
//		for v := range gMap {
//			if gMap1[v] > 1 {
//				return false
//			}
//		}
//	}
//	return true
//}
//
//func largerNumber(index, target int, nums []int) []int {
//	var items []int
//	for i := 0; i < index; i++ {
//		if target < nums[i] {
//			items = append(items, nums[i])
//		}
//	}
//	return items
//}

//func checkPossibility(nums []int) bool {
//	var length = len(nums)
//	if length == 1 {
//		return true
//	}
//	var count int
//	for index, value := range nums {
//		var number = largerNumber(index, value, nums)
//		if number >= 1 {
//			count++
//			if count > 1 {
//				return false
//			}
//		}
//	}
//	return true
//}
//
//func largerNumber(index, target int, nums []int) int {
//	var count int
//	for i := 0; i < index; i++ {
//		if target < nums[i] {
//			count++
//		}
//	}
//	return count
//}


//网上的解法
func checkPossibility(nums []int) bool {
	var length = len(nums)
	if length <= 2 {
		return true
	}
	var count int
	for i := 1; i < length; i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		count++
		if i >= 2 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return count <= 1
}
