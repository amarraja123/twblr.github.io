package control_structures

import (
	"math"
	"fmt"
)

// Can do without library.
func fizzBuzz(i int32) string {

	if((math.Mod(float64(i), float64(5)) == float64(0)) && (math.Mod(float64(i), float64(3)) == float64(0))){
		return "FizzBuzz"
	} else	if(math.Mod(float64(i), float64(2)) == float64(0)){
		return fmt.Sprintf("%v",i)
	} else	if(math.Mod(float64(i), float64(3)) == float64(0)){
		return "Fizz"
	} else	if(math.Mod(float64(i), float64(5)) == float64(0)){
		return "Buzz"
	}
	return ""
}
