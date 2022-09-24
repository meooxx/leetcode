package matchString

import "testing"

func TestMatchStr(t *testing.T) {
	a:= "   2123"
	b:= "-1232"
	c:="words ss 123"

	v1:=MatchStr(a)
	v2:=MatchStr(b)
	v3:=MatchStr(c)
	if v1 != 2123 || v2 != -1232 || v3 !=0 {
		t.Errorf("error occus")
	}


}