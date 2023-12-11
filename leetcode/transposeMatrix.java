class Solution {
  public int[][] transpose(int[][] nums) {
    int[][] ans = new int[nums[0].length][nums.length];

    for (int i = 0; i < nums.length; ++i)
      for (int j = 0; j < nums[0].length; ++j)
        ans[j][i] = nums[i][j];

    return ans;
  }
}
