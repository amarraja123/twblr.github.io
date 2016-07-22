package control_structures

func collatzChainLength(n int) int {
	var i =0;
	for ; n != 1; i++ {
		if (n %2 == 0){
			n = n/2
		} else {
			n = ((3 * n) + 1 )
		}
	}
	return i;
}
