package minus

import "testing"

func Test_minus(t *testing.T) {

	if Minus(5,2) != 4 {
		t.Fatal("Test_plus fail ")
	}

}
