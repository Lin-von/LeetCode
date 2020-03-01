package main

//func dfs(i int, edge [][]int, cur, walk []bool) bool {
//	if cur[i] {
//		return false
//	}
//
//	if walk[i] {
//		return true
//	}
//
//	cur[i] = true
//	walk[i] = true
//	for _, next := range edge[i] {
//		if !dfs(next, edge, cur, walk) {
//			return false
//		}
//	}
//	cur[i] = false
//
//	return true
//}
//
//func canFinish(numCourses int, prerequisites [][]int) bool {
//	edge := make([][]int, numCourses)
//	for i := 0; i < numCourses; i++ {
//		edge[i] = make([]int, 0)
//	}
//	for _, e := range prerequisites {
//		edge[e[1]] = append(edge[e[1]], e[0])
//	}
//	walked := make([]bool , numCourses)
//
//	for i := 0; i < numCourses; i++ {
//		if walked[i] {
//			continue
//		}
//
//		cur := make([]bool, numCourses)
//		if !dfs(i, edge, cur, walked) {
//			return false
//		}
//	}
//	return true
//}

// 入度表
func canFinish(numCourses int, prerequisites [][]int) bool {
	d := make([]int, numCourses)
	edge := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		edge[i] = make([]int, 0)
	}
	for _, e := range prerequisites {
		d[e[0]] ++
		edge[e[1]] = append(edge[e[1]], e[0])
	}
	
	queue := []int{}
	
	for k, v := range d {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	
	for len(queue) != 0 {
		for _, next := range edge[queue[0]] {
			d[next] --
			if d[next] == 0 {
				queue = append(queue, next)
			}
		}
		queue = queue[1:]
		numCourses --
	}
	if numCourses == 0 {
		return true
	} else {
		return false
	}
}