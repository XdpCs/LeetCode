package leetcode

// @Title        1486.XOR-Operation-in-an-Array.go
// @Description  1486.XOR-Operation-in-an-Array solution
// @Create       XdpCs 2023-09-25 13:32
// @Update       XdpCs 2023-09-25 13:32

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + i*2
	}
	return ans
}
