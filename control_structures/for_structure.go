package control_structures

func collatzChainLength(n int) int {
	var num = 0;
	for i := 0; n != 1; i++ {
		if (n %2 == 0){
			n = n/2
		} else {
			n = ((3 * n) + 1 )
		}
		num ++;
	}
	return num;
}
