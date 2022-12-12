class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        startSubstring  = 0
        endSubstring  = 0
        maxLength = 0
        unique_characters = set()
        while endSubstring  < len(s):
            if s[endSubstring] not in unique_characters:
                unique_characters.add(s[endSubstring ])
                endSubstring  += 1
                maxLength = max(maxLength, len(unique_characters))
            else:
                unique_characters.remove(s[startSubstring])
                startSubstring  += 1
        return maxLength
