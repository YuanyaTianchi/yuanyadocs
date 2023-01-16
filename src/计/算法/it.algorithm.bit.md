



题目：[136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/)

```go
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
```



题目：[191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)

```go
func hammingWeight(num uint32) int {
	counter := 0
	for num != 0 {
		num &= num - 1
		counter++
	}
	return counter
}
```

