package main

func dfs(s string, m map[string]map[string]float64, tar string, walk map[string]bool)  {
	for k, _ := range m[s] {
		if !walk[k] {
			walk[k] = true
			if _, ok := m[k][tar]; !ok {
				m[k][tar] = m[k][s] * m[s][tar]
				m[tar][k] = 1 / m[k][tar]
			}
			dfs(k, m, tar, walk)
		}
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := map[string]map[string]float64{}
	for i := 0; i < len(values); i++ {
		p := equations[i][0]
		q := equations[i][1]
		if _, ok := m[p]; !ok {
			m[p] = map[string]float64{}
			m[p][p] = 1.0
		}
		if _, ok := m[q]; !ok {
			m[q] = map[string]float64{}
			m[q][q] = 1.0
		}

		m[p][q] = values[i]
		m[q][p] = 1 / values[i]
	}

	for node, _ := range m {
		walk := map[string]bool{}
		walk[node] = true
		dfs(node, m, node, walk)
	}

	ret := []float64{}

	for i := 0; i < len(queries); i++ {
		p := queries[i][0]
		q := queries[i][1]
		if _, exist := m[p][q]; exist {
			ret = append(ret, m[p][q])
		} else {
			ret = append(ret, -1.0)
		}
	}
	return ret
}
