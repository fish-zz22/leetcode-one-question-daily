using System;

public class Solution
{
    public int ComputeArea(int ax1, int ay1, int ax2, int ay2, int bx1, int by1, int bx2, int by2)
    {
        // 右上角的点大于左下角的点
        var area1 = (ax2 - ax1) * (ay2 - ay1);
        var area2 = (bx2 - bx1) * (by2 - by1);
        var overlapWidth = Math.Min(ax2, bx2) - Math.Max(ax1, bx1);
        var overlapHeight = Math.Min(ay2, by2) - Math.Max(ay1, by1);
        var overlapArea = Math.Max(overlapWidth, 0) * Math.Max(overlapHeight, 0);
        return area1 + area2 - overlapArea;
    }
}