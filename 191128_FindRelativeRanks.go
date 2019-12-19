package main

import (
	"sort"
	"strconv"
)

func main() {
	var nums = []int{5, 1, 4, 2, 3}
	var answers = findRelativeRanks(nums)
	for _, answer := range answers {
		println(answer)
	}
}

func findRelativeRanks(nums []int) []string {

	var orders = make(map[int]int, 0)
	var length = len(nums)

	for i := 0; i < length; i++ {
		orders[i+1] = nums[i]
	}
	sort.Ints(nums)

	answer := make([]string, length)
	var order = 1
	var gmap = mapConvert(orders)
	for i := length - 1; i >= 0; i-- {
		//var key = getValueMap(orders, nums[i]) - 1
		var key = gmap[nums[i]] - 1
		switch order {
		case 1:
			answer[key] = "Gold Medal"
			break
		case 2:
			answer[key] = "Silver Medal"
			break
		case 3:
			answer[key] = "Bronze Medal"
			break
		default:
			answer[key] = strconv.Itoa(order)
		}
		order++
	}

	var answers = make([]string, 0)
	for i := 0; i < length; i++ {
		answers = append(answers, answer[i])
	}

	return answers
}

//自主实现由value获取key的方法
//func getValueMap(gmaps map[int]int, target int) int {
//	if gmaps == nil {
//		return -1
//	}
//	for k, v := range gmaps {
//		if v == target {
//			return k
//		}
//	}
//	return -1
//}

func mapConvert(gmap map[int]int) map[int]int {
	gmap1 := make(map[int]int, 0)
	for k, v := range gmap {
		gmap1[v] = k
	}
	return gmap1
}
