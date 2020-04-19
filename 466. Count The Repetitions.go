package main

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	l1 := len(s1)
	l2 := len(s2)

	if l1 == 0 || l2 == 0 || l1 * n1 < l2 * n2 {
		return 0
	}

	p1, p2 := 0, 0
	m1 := map[int]int{}
	m2 := map[int]int{}

	ret := 0
	for p1/l1 < n1 {
		if p1 % l1 == l1 - 1 {
			if v, ok := m1[p2 % l2]; ok {
				cl := (p1 - v) / l1
				cn := (n1 - 1 - p1 / l1) / cl
				s2num := p2 / l2 - m2[p2 % l2] / l2

				p1 += cl * cn * l1
				ret += cn * s2num
			} else {
				m1[p2 % l2] = p1
				m2[p2 % l2] = p2
			}
		}

		if s1[p1 % l1] == s2[p2 % l2] {
			if p2 % l2 == l2 - 1 {
				ret ++
			}
			p2 ++
		}
		p1 ++
	}
	return ret / n2
}
