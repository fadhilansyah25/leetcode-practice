// class Solution {
//     public int[] plusOne(int[] digits) {
//         if(digits[digits.length - 1] < 9) {
//             digits[digits.length - 1]++;
//             return digits;
//         }

//         ArrayList<Integer> res = new ArrayList<Integer>();

//         int carry = 1;

//         for(int i = digits.length - 1; i >= 0; i--) {
//             int sum  = carry + digits[i];

//             if(sum == 10) {
//                 res.add(0,0);
//             } else {
//                 res.add(0, digits[i] + carry);
//                 carry = 0;
//             }
//         }
//         if(carry != 0) res.add(0, 1);

//         // convert into new array from ArrayList
//         int[] result = new int[res.size()];

//         for(int i = 0; i < res.size(); i++) {
//             result[i] = res.get(i);
//         }

//         return result;
//     }
// }

class Solution {
    public int[] plusOne(int[] digits) {
       int n = digits.length;

       for(int i = n - 1; i >=0; i--) {
            if(digits[i] < 9) {
                digits[i]++;
                return digits;
            }
            digits[i] = 0;
       }

       int[] result = new int[n+1];
       result[0] = 1;
       return result;
    }
}
