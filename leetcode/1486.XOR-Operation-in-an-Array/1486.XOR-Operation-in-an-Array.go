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

func xorOperationMath(n int, start int) int {
	s := start >> 1
	e := start & n & 1
	return (sumXor(s-1)^sumXor(s+n-1))<<1 | e
}

func sumXor(x int) int {
	if x%4 == 0 {
		return x
	} else if x%4 == 1 {
		return 1
	} else if x%4 == 2 {
		return x + 1
	}
	return 0
}
