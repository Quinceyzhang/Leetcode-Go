package main

//输入一个按升序排序的整数数组（可能包含重复数字），你需要将它们分割成几个子序列，其中每个子序列至少包含三个连续整数。返回你是否能做出这样的分割？
//
// 
//
//示例 1：
//
//输入: [1,2,3,3,4,5]
//输出: True
//解释:
//你可以分割出这样两个连续子序列 :
//1, 2, 3
//3, 4, 5
//  
//
//示例 2：
//
//输入: [1,2,3,3,4,4,5,5]
//输出: True
//解释:
//你可以分割出这样两个连续子序列 :
//1, 2, 3, 4, 5
//3, 4, 5
// 
//
//示例 3：
//
//输入: [1,2,3,4,4,5]
//输出: False
// 
//
//提示：
//
//输入的数组长度范围为 [1, 10000]
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}
func isPossible(nums []int) bool {
	var length = len(nums)
	if length < 3 {
		return false
	}
	var gMap = make(map[int]int)
	for _, v := range nums {
		gMap[v]++
	}

	return true
}
