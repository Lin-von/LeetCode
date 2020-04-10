package main


func dfs(cur string, time int) string {
	ret := ""
	tmp := ""
	for i := 0; i < len(cur); i++ {
		if cur[i] > '0' && cur[i] <= '9' {
			t := 0
			for cur[i] >= '0' && cur[i] <= '9' {
				t *= 10
				t += int(cur[i] - '0')
				i ++
			}
			pos := i + 1
			cnt := 1
			for cnt > 0 {
				if cur[pos] == '[' {
					cnt ++
				} else if cur[pos] == ']' {
					cnt --
				}
				pos ++
			}
			pos --
			tmp += dfs(cur[i + 1:pos], t)
			i = pos
		} else {
			tmp += string(cur[i])
		}
	}

	for time > 0 {
		ret += tmp
		time --
	}
	return ret
}

func decodeString(s string) string {
	return dfs(s, 1)
}
