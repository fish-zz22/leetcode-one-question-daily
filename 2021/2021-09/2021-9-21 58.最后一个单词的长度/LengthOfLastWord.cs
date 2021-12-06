public class Solution {
    public int LengthOfLastWord(string s) 
    {
        var result = 0;
        for (var i = s.Length - 1; i >= 0; i--)
        {
            if ( s[i] == ' ')
            {
                if (result == 0)
                {
                    continue;
                }
                else
                {
                    break;
                }
            }
            result++;
        }
        return result;
    }
}