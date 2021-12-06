public class Solution
{
    public int[] Source { get; set; }
    private int[] Nums { get; set; }

    public int[] IndexArray { get; set; }

    public Solution(int[] nums)
    {
        Nums = nums;
        Source = new int[nums.Length];
        Array.Copy(nums,Source,nums.Length);
    }

    public int[] Reset()
    {
        Array.Copy(Source,Nums,Nums.Length);
        return Nums;
    }

    public int[] Shuffle()
    {
        var random = new Random();
        for (int i = 0; i < Nums.Length; i++)
        {
            var j = random.Next(Nums.Length);
            (Nums[i], Nums[j]) = (Nums[j], Nums[i]);
        }

        return Nums;
    }
}

/**
* Your Solution object will be instantiated and called as such:
* Solution obj = new Solution(nums);
* int[] param_1 = obj.Reset();
* int[] param_2 = obj.Shuffle();
*/