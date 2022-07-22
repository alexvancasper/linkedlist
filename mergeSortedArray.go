func mergeSortedArr(nums1, nums2 []int) (res []int) {
	if len(nums1) > len(nums2) {
		mergeSortedArr(nums2, nums1)
	}
	i := 0
	for j := 0; j < len(nums2); j++ {
		if i < len(nums1) {
			if nums1[i] < nums2[j] {
				res = append(res, nums1[i])
				i++
				j--
			} else if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				i++
			} else if nums1[i] > nums2[j] {
				res = append(res, nums2[j])
			}
		} else {
			res = append(res, nums2[j])
		}
	}
	return
}

