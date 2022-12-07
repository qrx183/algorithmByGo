package main

import (
	"math"
)

//无重复的最长字符串
func lengthOfLongestSubstring(s string) int {
	res := 0
	cacheMap := map[uint8]uint8{}
	l := 0
	r := 0

	n := len(s)
	for l < n {
		if l != 0 {
			delete(cacheMap, s[l-1])
		}
		ok := true
		if r < n {
			_, ok = cacheMap[s[r]]
		}
		for r < n && !ok {
			res = int(math.Max(float64(res), float64(r-l+1)))
			cacheMap[s[r]] = s[r]
			r++
			if r < n {
				_, ok = cacheMap[s[r]]
			}

		}
		l++
	}
	return res
}
