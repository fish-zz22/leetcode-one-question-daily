using System.Collections.Generic;
using System.Linq;

public class MapSum
{
    public Dictionary<string, int> Dict { get; set; } = new Dictionary<string, int>();

    public void Insert(string key, int val) => Dict[key] = val;

    public int Sum(string prefix) =>
        Dict.Where(x => x.Key.StartsWith(prefix))
            .Sum(x => x.Value);
}

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.Insert(key,val);
 * int param_2 = obj.Sum(prefix);
 */