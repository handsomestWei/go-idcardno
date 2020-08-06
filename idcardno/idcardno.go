package idcardno

import (
	"github.com/handsomestWei/go-idcardno/timex"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// 权重
var idNoWeightArray = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// 身份证号码
var idNoCheckCode = "10X98765432"

// 身份证号码正则匹配
var idCardNoRegexpPattern = "^([1-9]\\d{7}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])\\d{3})|([1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])((\\d{4})|\\d{3}[Xx]))$"

// 身份证信息
type IdCardNoInfo struct {
	IdCardNo    string // 身份证号码
	AreaCode    string // 地区编号
	AreaName    string // 地区名称
	BirthDayYMD string // 年月日生日，20060102
	Age         int    // 年龄
	Sex         int    // 性别，女0，男1
}

// 校验18位身份证号码有效性
func Validate18IdCardNo(idNo string) bool {
	if len(idNo) != 18 {
		return false
	}
	isMatch, _ := regexp.MatchString(idCardNoRegexpPattern, idNo)
	if !isMatch {
		return false
	}
	return getCheckDigit(idNo) == idNo[17:18]
}

// 自动生成18位身份证号码
func AutoCreate18IdCardNo() string {
	idNo := ""
	// 随机数种子
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	// 6位地区编码
	idNo += areaCodeList[ran.Intn(len(areaCodeList))]
	// 8位年月日生日
	idNo += timex.RandBirthDay().Format("20060102")
	// 2位顺序码
	idNo += strconv.Itoa(ran.Intn(9)+1) + strconv.Itoa(ran.Intn(10))
	// 1位性别，女双数，男单数
	idNo += strconv.Itoa(ran.Intn(10))
	// 1位校验位
	idNo += computerCheckDigit(idNo)
	return idNo
}

// 获取18位身份证号码信息
func Parse18IdCardNoInfo(idNo string) (bool, *IdCardNoInfo) {
	isIdCardNo := Validate18IdCardNo(idNo)
	if !isIdCardNo {
		return false, nil
	}
	return true, &IdCardNoInfo{
		IdCardNo:    idNo,
		AreaCode:    getAreaCode(idNo),
		AreaName:    getAreaName(idNo),
		BirthDayYMD: getBirthDayYMD(idNo),
		Age:         getAge(idNo),
		Sex:         getSex(idNo),
	}
}

// 获取校验位
func getCheckDigit(idNo string) string {
	data := idNo[0:17]
	s := 0
	for i, _ := range data {
		n, _ := strconv.Atoi(string(data[i]))
		s += n * idNoWeightArray[i]
	}
	y := s % 11
	return idNoCheckCode[y : y+1]
}

// 生成校验位
func computerCheckDigit(idNo string) string {
	checkSum := 0
	for i := 0; i < 17; i++ {
		n, _ := strconv.Atoi(string(idNo[i]))
		checkSum += ((1 << uint(17-i)) % 11) * n
	}
	checkDigit := (12 - (checkSum % 11)) % 11
	if checkDigit >= 10 {
		return "X"
	} else {
		return strconv.Itoa(checkDigit)
	}
}

func getAreaCode(idNo string) string {
	return idNo[0:6]
}

func getAreaName(idNo string) string {
	return areaNameMap[idNo[0:6]]
}

func getBirthDayYMD(idNo string) string {
	return idNo[6:14]
}

func getAge(idNo string) int {
	return timex.GetAgeByBirthDayYMD(idNo[6:14])
}

func getSex(idNo string) int {
	sex, _ := strconv.Atoi(idNo[16:17])
	return sex % 2
}
