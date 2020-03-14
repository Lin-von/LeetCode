package main

func leastInterval(tasks []byte, n int) int {
	hash := map[byte]int{}
	cnt := 0
	max := 0
	for _, s := range tasks {
		hash[s] ++
		if hash[s] > max {
			max = hash[s]
		}
	}
	for _, v := range hash {
		if v == max {
			cnt ++
		}
	}
	if (max - 1) * (n + 1) + cnt > len(tasks) {
		return (max - 1) * (n + 1) + cnt
	} else {
		return len(tasks)
	}
}
