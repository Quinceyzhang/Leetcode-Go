package main

//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//你找到的子数组应是最短的，请输出它的长度。
//
//示例 1:
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//var nums = []int{2, 3, 5, 3, 6, 5, 4, 6, 7, 8}
	//var nums = []int{2, 6, 4, 8, 10, 9, 15}
	//var nums = []int{6, 4}
	//var nums = []int{1, 2, 3, 3, 3}
	//var nums = []int{3, 3, 3, 3, 3}
	var nums = []int{1, 3, 2, 2, 2}

	println(findUnsortedSubarray(nums))
}
func findUnsortedSubarray(nums []int) int {
	var length = len(nums)
	var r, l = 0, 0

	var max = nums[0]
	var min = nums[length-1]
	for i := 0; i < length; i++ {
		if max > nums[i] {
			r = i
		} else {
			max = nums[i]
		}
		if min < nums[length-i-1] {
			l = length - i - 1
		} else {
			min = nums[length-i-1]
		}
	}
	if r == l {
		return 0
	} else {
		return r - l + 1
	}

}
