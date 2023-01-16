

+++

title = "array"
description = "it.algorithm.array"
tags = ["it", "algorithm"]

+++

# array



## 滑动窗口



### 最小覆盖子串

题目：[76. 最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)

```go
// 提取 t 中的字母，并且要判断是否包含字母，所以用 map 最合适
// 判断当前窗口是否满足要求：可以用一个 map 记录s的字符是否是t中的字符，且出现几次
func minWindow(s string, t string) string {
	// ts 记录 t 串中出现的字母及出现次数
	ts := make(map[uint8]int)
	for i, _ := range t {
		ts[t[i]] ++
	}
	// ss 记录 s 串中出现的且存在于 t 串中的字母及出现次数
	ss := make(map[uint8]int)

	// l 和 r 记录滑窗左右边界，counter 记录 s 串已满足出现次数的字符的个数，则当 counter 等于 ts 的大小时，表示找到一个此题的解
	l, r, counter := 0, 0, 0
	// 用于记录最优解变量，当然实际上只用 resL 和 minLin 也可以
	resL, resR, minLen := l, r, len(s)
	for r < len(s) {
		// 遍历 s 串，扩大右边界
		cr := s[r]
		// 如果 t 串包含字符，则应使 ss 中对应字符出现次数 +1
		if ts[cr] > 0 {
			ss[cr]++
			// 如果 s 串中字符出现次数 == t 串中该字符出现次数，表示该字符已满足出现次数的要求，则 count++
			if ss[cr] == ts[cr] {
				counter++
			}
		}
		r++

		// 表示找到一个满足需求的解，之后需要缩小窗口找到当前最优解
		for counter == len(ts) {
			cl := s[l]
			// 要去除的字符属于t串包含的字符，则需要使 ss 中对应--
			if ts[cl] > 0 {
				ss[cl]--
				// 当 s 串中字符出现次数 < t 串中该字符出现次数发生时，表示刚好不再满足需求，则[l,r)为当前轮的最优解
				if ss[cl] < ts[cl] {
					// 对比全局最优解判断是否替换
					length := r - l
					if length <= minLen {
						resL, resR, minLen = l, r, length
					}
					counter--
				}
			}
			l++
		}
	}
	return s[resL:resR]
}
```





题目：[567. 字符串的排列](https://leetcode-cn.com/problems/permutation-in-string/)

```go
func checkInclusion(s1 string, s2 string) bool {
	s1cs := make(map[uint8]int)
	for i, _ := range s1 {
		s1cs[s1[i]]++
	}
	s2cs := make(map[uint8]int)

	l, r, counter := 0, 0, 0
	for r < len(s2) {
		cr := s2[r]
		if s1cs[cr] > 0 {
			s2cs[cr]++
			if s2cs[cr] == s1cs[cr] {
				counter++
			}
		}
		r++

		// 因为 s1 的排列是固定长度的，因此对于 s2 满足要求的子串也是固定长度，缩小窗口时就无需循环了
		if r-l == len(s1) {
			if counter == len(s1cs) {
				return true
			}
			cl := s2[l]
			if s1cs[cl] > 0 {
				if s2cs[cl] == s1cs[cl] {
					counter--
				}
				s2cs[cl]--
			}
			l++
		}
	}
	return false
}
```



### 找到字符串中所有字母异位词

题目：[438. 找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

```go
// 满足条件是 s 子串与 p 长度相等，各字符出现次数相等
func findAnagrams(s string, p string) []int {
	ps := make(map[uint8]int)
	for i, _ := range p {
		ps[p[i]]++
	}
	ss := make(map[uint8]int)

	res := make([]int, 0)
	l, r, counter := 0, 0, 0
	for r < len(s) {
		// 扩
		rc := s[r]
		if ps[rc] > 0 {
			ss[rc]++
			if ss[rc] == ps[rc] {
				counter++
			}
		}
		r++

		if r-l == len(p) {
			if counter == len(ps) {
				res = append(res, l)
			}

			// 缩
			lc := s[l]
			if ps[lc] > 0 {
				if ss[lc] == ps[lc] {
					counter--
				}
				// --后置
				ss[lc]--
			}
			l++
		}
	}
	return res
}
```



### 无重复字符的最长子串

题目：[3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

```go
func lengthOfLongestSubstring(s string) int {
	ss := make(map[uint8]int)

	l, r, counter := 0, 0, 0
	res := counter
	for r < len(s) {
		rc := s[r]
		ss[rc]++
		counter++
		r++

		for ss[rc] > 1 {
			lc := s[l]
			ss[lc]--
			counter--
			l++
		}

		if counter > res {
			res = counter
		}
	}
	return res
}
```



