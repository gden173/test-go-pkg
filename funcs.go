package funcs


// Function which squares an input array of numbers 
func SquareNum(nums []float32) []float32 {
	for i := 0; i < len(nums); i ++ {
		nums[i] *= nums[i]
	}
	return nums
}
