package main

func postTest(l, r int, order []int) bool {
	if l >= r {
		return true
	}

	tmp := order[r]
	pos := l
	for order[pos] < tmp {
		pos ++
	}
	cpos := pos
	for pos < r {
		if order[pos] < tmp {
			return false
		}
		pos ++
	}
	return postTest(l, cpos - 1, order) && postTest(cpos, r - 1, order)

}

func verifyPostorder(postorder []int) bool {
	return postTest(0, len(postorder) - 1, postorder)
}
