class Solution {
    public int climbStairs(int n) {
        int prev = 0;
        int current = 1;

        for(int i = 0; i < n; i++) {
            int temp = current;
            current += prev;
            prev = temp;
        }


        return current;
    }
}