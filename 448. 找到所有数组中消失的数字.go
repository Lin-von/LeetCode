package main

func Abs(num int) int {
	if num < 0 {
		num = 0 - num
	}
	return num
}


func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		nums[Abs(num) - 1] = 0 - Abs(nums[Abs(num) - 1])
	}

	arr := []int{}
	for i := 1; i <= len(nums); i++ {
		if nums[i - 1] > 0 {
			arr = append(arr, i)
		}
	}
	return arr
}
