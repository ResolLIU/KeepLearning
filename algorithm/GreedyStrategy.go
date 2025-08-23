package algorithm

import "sort"

// []int{} 求平均分配后最小极差，最多count次操作
// 例如 [1,5,5,6,8] count = 3
// [4,5,5,5,6] -> return 2

func Distribute(nums []int, count int) int {
	// 朴素的想，每一次将大的移到小的地方去
	cnt := make(map[int]int)
	sum, n := 0, len(nums)
	for _, i := range nums {
		cnt[i]++
		sum += i
	}
	var dic []int
	for i := range cnt {
		dic = append(dic, i)
	}
	ave := sum / n
	avh := ave
	if sum%n != 0 {
		avh++
	} // 要考虑能否正好分配，若不能整除，则最多优化到极差为1
	sort.Ints(dic)
	l, r := 0, len(dic)-1 // 最大和最小值下标
	// 贪心的想，大的先消耗count，最多变为平均数，可以证明，大的最多变为平均数
	high, low := dic[r], dic[l]
	c := count
	for high > avh && r-1 >= 0 {
		target := max(avh, dic[r-1]) // 把最大的一批先变小
		cost := cnt[high] * (high - target)
		if c-cost >= 0 {
			c -= cost
			h := cnt[high]
			high = target
			r--
			cnt[target] += h
		} else { // 不够操作了
			high -= c / cnt[high]
			c = 0
			break
		}
	}
	if c > 0 {
		return avh - ave
	}
	c = count
	for low < ave && l+1 < len(dic) {
		target := min(ave, dic[l+1])
		cost := cnt[low] * (target - low)
		if c-cost >= 0 {
			c -= cost
			lo := cnt[low]
			low = target
			l++
			cnt[target] += lo
		} else {
			low += c / cnt[low]
			break
		}
	}
	return high - low
}
