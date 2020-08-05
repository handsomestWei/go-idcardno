package idcardno

import "testing"

func TestValidate18IdCardNo_1(t *testing.T) {
	idNo := "110101199003074119"
	if !Validate18IdCardNo(idNo) {
		t.Fail()
	}
}

func TestValidate18IdCardNo_2(t *testing.T) {
	idNo := "11010119900307411X"
	if !Validate18IdCardNo(idNo) {
		t.Fail()
	}
}