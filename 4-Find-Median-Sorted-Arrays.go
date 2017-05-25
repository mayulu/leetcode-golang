package leetcode_go_mayulu

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	l := (m + n + 1) / 2
	r := (m + n + 2) / 2

	return float64(find_kth(nums1, m, nums2, n, l)+find_kth(nums1, m, nums2, n, r)) / 2.0

}

func find_kth(ss []int, m int, ls []int, n int, k int) int {

	if m > n {
		find_kth(ls, n, ss, m, k)
	}
	if m == 0 {
		return ls[k-1]
	}
	if k == 1 {
		return min(ss[0], ls[0])
	}

	i := min(m, k/2)
	j := min(n, k/2)

	if ss[i-1] > ls[j-1] {
		return find_kth(ss, m, ls[j:], n-j, k-j)
	} else {
		return find_kth(ss[i:], m-i, ls, n, k-i)
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
