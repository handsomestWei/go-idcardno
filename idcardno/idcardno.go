package idcardno

import (
	"regexp"
	"strconv"
)

var idNoWeightArray = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var idNoCheckCode = "10X98765432"

// 校验18位身份证号码有效性
// @see https://www.cnblogs.com/xuejianbest/p/10285176.html
func Validate18IdCardNo(idNo string) bool {

	if len(idNo) != 18 {
		return false
	}

	isMatch, _ := regexp.MatchString("^([1-9]\\d{7}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])\\d{3})|([1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])((\\d{4})|\\d{3}[Xx]))$", idNo)
	if !isMatch {
		return false
	}

	data := idNo[0:17]
	s := 0
	for i, _ := range data {
		n, _ := strconv.Atoi(string(data[i]))
		s += n * idNoWeightArray[i]
	}
	y := s % 11
	code := idNoCheckCode[y : y+1]
	value := idNo[17:18]
	return code == value
}
