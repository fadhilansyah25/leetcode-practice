func searchInsert(nums []int, target int) int {
    var pos int = 0;

    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            pos = i
            break;
        }

        if nums[i] < target {
            pos++;
        }
    }

    return pos
}