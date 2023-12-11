import java.util.Arrays;

class Solution {
    public int[] decrypt(int[] code, int k) {
        int n = code.length;
        int[] result = new int[n];

        if (k > 0) {
            for (int i = 0; i < n; i++) {
                result[i] = calculateSumOfKNextNums(code, i, k);
            }
        } else if (k < 0) {
            for (int i = 0; i < n; i++) {
                result[i] = calculateSumOfKPreviousNums(code, i, k);
            }
        }

        return result;
    }

    public int calculateSumOfKNextNums(int[] code, int start, int k) {
        int result = 0;
        int n = code.length;

        for (int j = 1; j <= k; j++) {
            result += code[(start + j) % n];
        }

        return result;
    }

    public int calculateSumOfKPreviousNums(int[] code, int start, int k) {
        int result = 0;
        int n = code.length;

        for (int j = 1; j <= -k; j++) {
            result += code[(start - j + n) % n];
        }

        return result;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] code = {1, 2, 3, 4};
        int k = 2;
        int[] result = solution.decrypt(code, k);
        System.out.println(Arrays.toString(result));
    }
}
