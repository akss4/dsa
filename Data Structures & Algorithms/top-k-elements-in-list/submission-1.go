func topKFrequent(nums []int, k int) []int {
	frequent :=  make(map[int]int)
	result := []int{}
	for _ , value := range nums{
		frequent[value]++
	}
	for i := 0; i<k; i++{
	maxfreq := 0
	maxvalue := 0
	
	 for value , count := range frequent{
		if count > maxfreq{
			maxfreq = count
			maxvalue = value 
		}
		
	}
	result = append(result , maxvalue)
	delete(frequent, maxvalue)
}
return result
}

