package revert

import "testing"

func TestRevertNum(t *testing.T){
	v1, v2, v3 := 120, 213,321
	t1:= RevertNum(v1)
	t2:=RevertNum(v2)
	t3:= RevertNum(v3)

	if t1 != 21 || t2!= 321 || t3 != 123 {
		t.Errorf("origin values 1:%d=>%d,2:%d=>%d, 3:%d>%d" , v1, v2,v3, t1,t2,t3)
	}

	if RevertNum(-21) != -12 {
		t.Errorf("origin -21 => %d", RevertNum(-21))
	}

}