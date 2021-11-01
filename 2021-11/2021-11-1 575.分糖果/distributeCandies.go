package main

func main() {

}

func distributeCandies(candyType []int) int {
	// 三种情况：
	// 1.糖果种类数量大于数组长的一半  此时妹妹拿到的糖果种类为数组长的一半
	// 2.糖果种类数量等于数组长的一半  此时妹妹拿到的糖果种类为数组长的一半
	// 3.糖果种类数量小于数组长的一半  此时妹妹拿到的糖果种类为糖果种类的数量

	// 映射空struct不占用内存
	candyMap := make(map[int]struct{}, len(candyType))
	l := len(candyType) / 2
	for _, v := range candyType {
		if _, ok := candyMap[v]; ok {
			continue
		}
		candyMap[v] = struct{}{}
	}

	if len(candyMap) >= l {
		return l
	}
	return len(candyMap)
}
