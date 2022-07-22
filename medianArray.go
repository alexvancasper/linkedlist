func medianArray(nums1, nums2 []int) float64 {
	var res []int
	i := 0
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for j := 0; j < len(nums2); j++ {
		if i < len(nums1) {
			if nums1[i] < nums2[j] {
				res = append(res, nums1[i])
				i++
				j--
			} else if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				res = append(res, nums2[j])
				i++
			} else if nums1[i] > nums2[j] {
				res = append(res, nums2[j])
			}
		} else {
			res = append(res, nums2[j])
		}
	}
	for l := i; l < len(nums1); l++ {

		res = append(res, nums1[l])
	}

	fmt.Println(res)

	modRes := len(res) % 2
	central := int(len(res) / 2)
	if modRes != 0 {
		return float64(res[int(central)])
	}
	return float64(res[central-1]+res[central]) / 2

}
