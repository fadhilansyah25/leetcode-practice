class Solution {
    public String addBinary(String a, String b) {
        int i = a.length() - 1;
        int j = b.length() - 1;
        int carry = 0;

        String res = "";

        while(i >= 0 || j >= 0 ) {
            int sum = 0;
            sum += carry;

            if (i >= 0) sum += a.charAt(i) - '0';

            if (j >= 0) sum += b.charAt(j) - '0';

            carry = sum > 1 ? 1 : 0;
            res = (sum % 2) + res;

            i--;
            j--;
        }


        return carry != 0 ? carry + res : res;
    }
}