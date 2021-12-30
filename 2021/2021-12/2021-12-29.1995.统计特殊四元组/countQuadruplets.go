package main

func main() {

}

func countQuadruplets(nums []int) (res int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						res++
					}
				}
			}
		}
	}
	return
}
