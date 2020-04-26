package main

var parent map[string]string

func find(s string) string {
	if s != parent[s] {
		parent[s] = find(parent[s])
	}
	return parent[s]
}

func union(a, b string)  {
	p := find(a)
	q := find(b)
	if p != q {
		parent[p] = q
	}
}

func isSimilar(a, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count ++
		}
	}
	return count <= 2
}

func scan(s string, arr map[string]bool)  {
	tmp := []byte(s)
	for i := 0; i < len(s) - 1; i ++ {
		for j := 0; j < len(s); j ++ {
			tmp[i], tmp[j] = tmp[j], tmp[i]
			if _, ok := arr[string(tmp)]; ok {
				union(string(tmp), s)
			}
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
	}
}

func numSimilarGroups(A []string) int {
	ret := 0
	n := len(A)
	l := len(A[0])
	parent = map[string]string{}
	hash := map[string]bool{}

	for _, s := range A {
		hash[s] = true
		parent[s] = s
	}

	if n < l * l {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if isSimilar(A[i], A[j]) {
					union(A[i], A[j])
				}
			}
		}
	} else {
		for i := 0; i < n; i++ {
			scan(A[i], hash)
		}
	}

	for k, v := range parent {
		if k == v {
			ret ++
		}
	}

	return ret

}
