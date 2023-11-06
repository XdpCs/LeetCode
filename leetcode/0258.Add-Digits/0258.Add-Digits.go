package leetcode

// @Title        0258.Add-Digits.go
// @Description  0258.Add-Digits solution
// @Create       XdpCs 2023-11-06 12:23
// @Update       XdpCs 2023-11-06 12:23

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}
