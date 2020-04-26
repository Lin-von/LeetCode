package main

func entityParser(text string) string {
	d := map[string]string{
		"&quot;":"\"",
		"&apos;":"'",
		"&amp;":"&",
		"&gt;":">",
		"&lt;":"<",
		"&frasl;":"/",
	}

	ret := ""
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			found := false
			for k, v := range d {
				if i + len(k) <= len(text) && text[i:i + len(k)] == k {
					ret += v
					i += len(k) - 1
					found = true
					break
				}
			}
			if !found {
				ret += string(text[i])
			}
		} else {
			ret += string(text[i])
		}

	}
	return ret
}
