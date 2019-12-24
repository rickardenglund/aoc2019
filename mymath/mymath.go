package intmath

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Exp2(x int) int {
	return 1 << x
}
