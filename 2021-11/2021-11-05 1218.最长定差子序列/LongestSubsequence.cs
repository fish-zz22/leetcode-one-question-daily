public class Solution {
    public int LongestSubsequence(int[] arr, int difference) {
        var result = 0;
        // dp[i]
        var dp = new Dictionary<int, int>();
        foreach (var i in arr)
        {
            if (dp.TryGetValue(i - difference, out var pre))
            {
                dp[i] = pre + 1;
            }
            else
            {
                dp[i] = 1;
            }
            result = Math.Max(result,dp[i]);
        }

        return result;
    }
}