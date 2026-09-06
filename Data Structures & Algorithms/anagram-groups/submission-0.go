func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _ , word := range strs{
		var  fingerprint [26]int
		for _ , char := range word {
			fingerprint[char-'a']++
		}
		groups[fingerprint]= append(groups[fingerprint],word)

	}
	result := [][]string{}
	for _ , group := range groups{
		result = append(result ,  group)
	} 
	return result

}
