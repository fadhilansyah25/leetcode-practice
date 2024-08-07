func removeElement(nums []int, val int) int {
	// nums = append(nums[:1], nums[2:]...)
	// for i, v := range nums {
	// 	if v == val {
	// 		nums = append(nums[:i], nums[i+1:]...)
	// 	}

	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	fmt.Println(nums)

	return len(nums)
}