package reverseStr

import "testing"


func TestReverse(t *testing.T) {
	v := Reverse("1321223")
	if v != "3221231" {
		t.Errorf("reverse error should: %s != result: %s", "321231", v)
	}
	v = Reverse("abcd")
	if v != "dcba" {
		t.Errorf("reverse error should: %s != result: %s", "abcd", v)
	}

}