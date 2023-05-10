package bnrsrch

func GetTargetIndex(arr *[]int, target int) int {
	l := 0
	r := len(*arr)
	for l <= r {
		mid := (l + r) / 2
		if (*arr)[mid] == target {
			return mid
		} else if (*arr)[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
