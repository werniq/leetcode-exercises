class Solution {
    public int maxArea(int[] height) {
        int maxArea = 0;
        int left = 0;
        int right = height.length - 1;

        while (left < right) {
            // width of the container
            int width = right - left;

            // height of the container
            int h = Math.min(height[left], height[right]);

            // calc area
            int area = width * h;
            maxArea = Math.max(maxArea, area);

            // move the pointers based on the height
            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }

        return maxArea;
    }
}
