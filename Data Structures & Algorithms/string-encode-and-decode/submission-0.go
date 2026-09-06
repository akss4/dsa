type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded:=""
	for _ , str := range strs{

	encoded += strconv.Itoa(len(str)) + "#" + str
		
	}
return encoded
}

func (s *Solution) Decode(encoded string) []string {
	
	result := []string{}
	i := 0
	for i < len(encoded){
		length := 0
		for encoded[i] != '#'{ // remove #
			length = length*10 + int(encoded[i]-'0') // we are also checking dual digit number 
			i++


		}
		i++
	    	result =append(result, encoded[i:i+length])
			i = i+length

	}
	return result

}
