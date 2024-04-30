package leetcode

// @Title        1512.Number-of-Good-Pairs.go
// @Description  1512.Number-of-Good-Pairs solution
// @Create       XdpCs 2023-09-25 16:39
// @Update       XdpCs 2023-09-25 16:39

func numIdenticalPairs(nums []int) int {
	hashMap := map[int]int{}
	for _, n := range nums {
		hashMap[n]++
	}
	count := 0
	for _, v := range hashMap {
		count += v * (v - 1) / 2
	}
	return count
}
