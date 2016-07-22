package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {

	result := []int32{}
	for _, eachVal := range vals{
		 j:= eachVal * eachVal
		result = append(result, j);
	}
	return result

}

func filterInts(op filterOperation, vals []int32) []int32 {
	result := []int32{}
	for _, eachVal := range vals{
		if(eachVal > 2){
			result = append(result, eachVal);
		}
	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {

	result := []string{}
	for _, eachVal := range dest{
		result = append(result, eachVal);
	}
	for _, eachVal := range newValues{
		result = append(result, eachVal);
	}
	return result
}

func equals(list1 []string, list2 []string) bool {

	if list1 == nil && list2 == nil {
		return true;
	}

	if list1 == nil || list2 == nil {
		return false;
	}

	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func partialReverse(src []int, from, to int) []int {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}

	return src
}
