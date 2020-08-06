package idcardno

import (
	"fmt"
	"testing"
)

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

func TestParse18IdCardNoInfo(t *testing.T) {
	idNo := "110101199003074119"
	result, info := Parse18IdCardNoInfo(idNo)
	if !result {
		t.Fail()
	} else {
		fmt.Println(fmt.Sprintf("%v", *info))
	}
}

func TestAutoCreate18IdCardNo(t *testing.T) {
	idNo := AutoCreate18IdCardNo()
	fmt.Println(idNo)
	if !Validate18IdCardNo(idNo) {
		t.Fail()
	}
	result, info := Parse18IdCardNoInfo(idNo)
	if !result {
		t.Fail()
	} else {
		fmt.Println(fmt.Sprintf("%v", *info))
	}
}
