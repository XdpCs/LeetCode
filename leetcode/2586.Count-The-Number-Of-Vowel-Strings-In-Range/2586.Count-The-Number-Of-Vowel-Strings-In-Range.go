package leetcode

// @Title        2586.Count-The-Number-Of-Vowel-Strings-In-Range.go
// @Description  2586.Count-The-Number-Of-Vowel-Strings-In-Range solution
// @Create       XdpCs 2023-11-23 10:48
// @Update       XdpCs 2023-11-23 10:48

func vowelStrings(words []string, left int, right int) int {
	hashMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	count := 0

	for i := left; i <= right; i++ {
		if hashMap[words[i][0]] && hashMap[words[i][len(words[i])-1]] {
			count++
		}
	}
	return count
}
