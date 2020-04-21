package main


func numberOfSubarrays(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	cnt := 0
	l := 0
	r := 0
	ret := 0

	for r < len(nums) {
		for cnt < k && r < len(nums) {
			if nums[r] & 1 == 1 {
				cnt ++
			}
			r ++
		}
		posr := r
		for r < len(nums) && nums[r] & 1 == 0 {
			r ++
		}

		posl := l
		for cnt == k {
			if nums[l] & 1 == 1 {
				cnt --
			}
			l ++
		}
		ret += (l - posl) * (r - posr + 1)
	}
	return ret
}
