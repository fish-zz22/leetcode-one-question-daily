public class Solution {
    public int FindNumberOfLIS(int[] nums) {
        var result = 0;
        var maxLength = 1;
        var dp = new int[nums.Length];
        var count = new int[nums.Length];

        for (int i = 0; i < nums.Length; i++)
        {
            dp[i] = 1;
            count[i] = 1;
            for (int j = 0; j < i; j++)
            {
                if (nums[i] > nums[j])
                {
                    if (dp[j] + 1 > dp[i])
                    {
                        dp[i] = dp[j] + 1;
                        count[i] = count[j]; // 重置计数
                    }
                    else if (dp[j] + 1 == dp[i])
                    {
                        count[i] += count[j];
                    }
                }
            }
        }
        // 搜索长度最大的数量和
        for (int i = 0; i < nums.Length; i++)
        {
            if (dp[i] > maxLength)
            {
                maxLength = dp[i];
                result = count[i]; // 重置计数
            }
            else if (dp[i] == maxLength)
            {
                result += count[i];
            }
        }

        return result;
    }
}