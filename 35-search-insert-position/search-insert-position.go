func searchInsert(nums []int, target int) int {
    i := 0;

    for _, value := range nums {
        if value >= target {
            break
        }

        i++;
    }

    return i;
}