func searchInsert(nums []int, target int) int {
    var pos int = 0;

    for pos < len(nums) {
        if nums[pos] == target {
            break;
        }

        if nums[pos] < target {
            pos++;
        } else {
            break
        }
    }

    return pos
}