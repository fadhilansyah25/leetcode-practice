func removeDuplicates(nums []int) int {
	temp := nums[0]
	count := 1

    fmt.Println(temp)

	for _, v := range nums {
		if v != temp {
			temp = v
            nums[count] = v
			count++
		}
	}

	return count
}