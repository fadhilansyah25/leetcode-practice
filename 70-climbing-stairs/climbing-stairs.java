class Solution {
    public int climbStairs(int n) {
        int prev = 0;
        int current = 1;

        int result = 1;

        for(int i = 1; i < n; i++) {
            int temp = current;
            current += prev;
            prev = temp;

            result = current + prev;
        }


        return result;
    }
}