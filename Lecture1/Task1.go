package Lecture1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		match := target - nums[i]
		idx, ok := m[match]
		if ok {
			return []int{idx, i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}
