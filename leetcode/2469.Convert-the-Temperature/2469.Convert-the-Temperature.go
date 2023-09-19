package leetcode

// @Title        2469.Convert-the-Temperature.go
// @Description  2469.Convert-the-Temperature solution
// @Create       XdpCs 2023-09-18 21:05
// @Update       XdpCs 2023-09-18 21:05

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}
