package leetcode

// @Title        1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer.go
// @Description  1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer solution
// @Create       XdpCs 2023-11-06 12:23
// @Update       XdpCs 2023-11-06 12:23

func subtractProductAndSum(n int) int {
	sum := 0
	multiSum := 1
	for ; n != 0; n /= 10 {
		ans := n % 10
		sum += ans
		multiSum *= ans
	}

	return multiSum - sum
}
