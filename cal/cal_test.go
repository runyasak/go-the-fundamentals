package cal

import "testing"

func TestSumOfFirstThree(t *testing.T) {
	given := 2
	want := 6

	get := SumOfFirst(given)

	if want != get {
		t.Errorf("given %d want %d, but get %d\n", given, want, get)
	}
}