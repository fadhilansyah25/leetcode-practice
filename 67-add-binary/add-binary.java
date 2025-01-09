class Solution {
    public String addBinary(String a, String b) {
        int i = a.length() - 1;
        int j = b.length() - 1;
        int carry = 0;

        String res = "";

        while(i >= 0 || j >= 0 ) {
            int sum = 0;
            sum += carry;

            if (i >= 0) {
               sum += a.charAt(i) - '0';
            }

            if(j >= 0) {
                sum += b.charAt(j) - '0';
            }

            if(sum >= 2) {
                carry = 1;
            } else {
                carry = 0;
            }

            res = (sum % 2) + res;

            i--;
            j--;
        }

        if(carry > 0) res = carry + res;

        return res;
    }
}