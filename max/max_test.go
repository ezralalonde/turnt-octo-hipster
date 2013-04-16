package max

import (
	"testing"
)

func tDis(in1, in2 []int, out int, t *testing.T) {
	f := MaximumDistance
	ans := f(in1, in2)
	if ans != out {
		t.Errorf("MaximumDistance(%v, %v) = %v, wanted %v", in1, in2, ans, out)
	}
}

func TestDistance1(t *testing.T) {
	a := []int{8, 8, 4, 4, 4, 3, 3, 3, 1}
	b := []int{9, 9, 8, 8, 6, 5, 5, 4, 3}
	tDis(a, b, 5, t)
}
func TestDistance2(t *testing.T) {
	a := []int{6, 5, 4, 4, 4, 4, 4}
	b := []int{3, 3, 3, 3, 3, 3, 3}
	tDis(a, b, 0, t)
}
