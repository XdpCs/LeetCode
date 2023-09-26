package leetcode

// @Title        1512.Number-of-Good-Pairs.go
// @Description  1512.Number-of-Good-Pairs solution
// @Create       XdpCs 2023-09-25 16:39
// @Update       XdpCs 2023-09-25 16:39

func numIdenticalPairs(nums []int) int {
	n := len(nums)
	hashMap := make(map[int]int, n)
	sum := 0
	for _, num := range nums {
		hashMap[num]++
	}

	for _, v := range hashMap {
		if v > 1 {
			sum += v * (v - 1) / 2
		}
	}
	return sum
}
