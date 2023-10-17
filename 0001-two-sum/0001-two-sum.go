func twoSum(nums []int, target int) []int {
    output := []int{}

    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                output = append(output, i, j)
                break
            }
        }
    }

    return output
}