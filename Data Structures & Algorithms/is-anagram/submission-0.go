func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
freqs := make(map[rune]int)
for _ , char := range s{
	freqs[char]++
}
freqt := make(map[rune]int)
for _, char := range t{
	freqt[char]++
}
for char , count := range freqs{
	if count != freqt[char]{
		return false
	}
}
	return true
}
