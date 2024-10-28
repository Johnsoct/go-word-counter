func WordCount(s string) map[string]int {
	m = make(map[string]int)
	strs := strings.Fields(s)
	
	for _, word := range strs {
		wordValue, ok := m[word]
		
		if ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	
	return m
}
