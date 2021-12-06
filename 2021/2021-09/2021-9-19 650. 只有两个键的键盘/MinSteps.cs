public class Solution {
    public int MinSteps(int n) {
        // dp[1] = 0
        // dp[2] = 2
        // dp[3] = 3
        // dp[4] = 4
        var dp = new int[n+1];
        dp[0] = 0;
        dp[1] = 0;
        for(int i = 2; i <= n; i++)
        {
            dp[i] = int.MaxValue;
            for (int j = 1; j * j <= i; j++)
            {
                if(i % j == 0)
                {
                    dp[i] = Math.Min(dp[i], dp[j] + i / j);
                    dp[i] = Math.Min(dp[i], dp[i / j] + j);
                }
            }
        }

        return dp[n];
    }
}