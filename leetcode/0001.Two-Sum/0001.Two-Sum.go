package leetcode

// @Title        0001.Two-Sum.go
// @Description
// @Create       XdpCs 2024-05-05 17:05
// @Update       XdpCs 2024-05-05 17:05

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			} else if nums[i]+nums[j] > target {
				continue
			}
		}
	}
	return []int{0, 0}
}

func twoSumHash(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		if j, ok := hashMap[target-v]; ok {
			return []int{i, j}
		}
		hashMap[v] = i
	}
	return []int{0, 0}
}
