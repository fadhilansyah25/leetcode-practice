class Solution {
    public int[] plusOne(int[] digits) {
        if(digits[digits.length - 1] < 9) {
            digits[digits.length - 1]++;
            return digits;
        }

        ArrayList<Integer> res = new ArrayList<Integer>();

        int carry = 1;

        for(int i = digits.length - 1; i >= 0; i--) {
            int sum  = carry + digits[i];

            if(sum == 10) {
                res.add(0,0);
            } else {
                res.add(0, digits[i] + carry);
                carry = 0;
            }
        }

        // 999
        // 1, 0, 0, 0

        if(carry != 0) res.add(0, 1);

        int[] result = new int[res.size()];

        for(int i = 0; i < res.size(); i++) {
            result[i] = res.get(i);
        }

        return result;
    }
}