public class Solution {
    public int LengthOfLastWord(string s) {
        string[] list=s.Split(' ');
        int len=0;
        foreach(string a in list)
        {
            if(a.Length>0)
                len=a.Length;
        }
        return len;
    }
}
