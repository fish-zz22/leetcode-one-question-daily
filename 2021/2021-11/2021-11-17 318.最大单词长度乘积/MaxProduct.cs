using System.Collections.Generic;
using System.Linq;

public class Solution
{
    public int MaxProduct(string[] words)
    {
        int result = 0;
        var set = new HashSet<char>();
        for (var i = 0; i < words.Length; i++)
        {
            var currentWord = words[i];
            set.Clear();
            foreach (var c in currentWord)
            {
                set.Add(c);
            }

            for (var j = i + 1; j < words.Length; j++)
            {
                var nextWord = words[j];
                // 不可能又更大的值了
                if (currentWord.Length * nextWord.Length <= result)
                {
                    continue;
                }
                // 重复时跳过
                if (nextWord.Any(x=> set.Contains(x)))
                {
                    continue;
                }
                // 更新最大值
                result = currentWord.Length * nextWord.Length;
            }
        }

        return result;
    }
}
