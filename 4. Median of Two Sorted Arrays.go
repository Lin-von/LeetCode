package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l, r int
	var maxL, minR int

	m := len(nums1)
	n := len(nums2)

	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	l, r = 0, m
	tmp := (m + n + 1) / 2
	for l <= r {
		i := (l + r) / 2
		j := tmp - i
		if i < r && nums1[i] < nums2[j - 1] {
			l = i + 1
		} else if i > l && nums2[j] < nums1[i - 1] {
			r = i - 1
		} else {
			if i == 0 {
				maxL = nums2[j - 1]
			} else if j == 0 {
				maxL = nums1[i - 1]
			} else {
				if nums1[i - 1] > nums2[j - 1] {
					maxL = nums1[i - 1]
				} else {
					maxL = nums2[j - 1]
				}
			}
			if (m + n) & 1 == 1{
				return float64(maxL)
			}

			if i == m {
				minR = nums2[j]
			} else if j == n {
				minR = nums1[i]
			} else {
				if nums1[i] > nums2[j] {
					minR = nums2[j]
				} else {
					minR = nums1[i]
				}
			}
			return float64(maxL + minR) / 2
		}
 	}

	return 0

 }
