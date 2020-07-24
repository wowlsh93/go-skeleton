package plus

import "testing"

func Test_plus(t *testing.T) {

	if Plus(1,2) != 4 {
		t.Fatal("Test_plus fail ")
	}

}
