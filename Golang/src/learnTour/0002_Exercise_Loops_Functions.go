package learnTour

import "fmt"
import "math"

/*
Calc *sqrt()*.
function is:
	guess z
	loop z -= (z*z - x) / (2*z)
	return z
*/

func mysqrt(x float64) (z float64) {
	delta := 0.000000001
	z = 1.0
	times := 0
	for y := 0.0; (y-z > delta) || (z-y > delta); times++ {
		y = z
		z = y - (y*y-x)/(2*y)
	}
	fmt.Printf("mysqrt  Loop %d times:  ", times)
	return z
}

func mysqrt2(x float64) (z float64) {
	delta := 0.000000001
	z = x / 2
	times := 0
	for y := 0.0; (y-z > delta) || (z-y > delta); times++ {
		y = z
		z = y - (y*y-x)/(2*y)
	}
	fmt.Printf("mysqrt2 Loop %d times:  ", times)
	return z
}

func Main0002() {
	x := 564684685136543841.0
	fmt.Println("The Standard Answer is:", math.Sqrt(x))
	fmt.Println(mysqrt(x))
	fmt.Println(mysqrt2(x))
}
