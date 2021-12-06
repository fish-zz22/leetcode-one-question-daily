using System;
using System.Linq;

public class Solution
{
    public int FindMinMoves(int[] machines)
    {
        // 求和 找到平均数 
        // 如果不能平均就 return -1
        // 如果一个机子已经达到平均，就传送到其他洗衣机
        var sum = machines.Sum();
        if (sum % machines.Length != 0)
        {
            return -1;
        }

        var mean = sum / machines.Length;

        var result = 0;
        var count = 0;
        foreach (var machine in machines)
        {
            var temp = machine - mean;
            // 因为每次只能移动相邻的，所以要叠加
            count += temp;
            // 需要移入的和需要移除的里面找一个最大值
            result = Math.Max(result, Math.Max(Math.Abs(count), temp));
        }

        return result;
    }
}