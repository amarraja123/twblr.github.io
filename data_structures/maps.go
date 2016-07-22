package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	result := []string{}
	for key,value := range data {
		for _, val := range value{
			if(interest == val){
				result = append(result,key)
			}
		}


	}
	return result
}

func contains(src []string, elem string) bool {
	for _, val := range src{
		if(elem == val){
			return true
		}
	}
	return false
}
