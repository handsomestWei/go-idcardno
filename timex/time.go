package timex

import (
	"math/rand"
	"time"
)

// 随机生成生日
func RandBirthDay() time.Time {
	start := time.Date(time.Now().Year()-40, 1, 0, 0, 0, 0, 0, time.Local).Unix()
	end := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.Local).Unix()
	delta := end - start
	sec := rand.New(rand.NewSource(time.Now().Unix())).Int63n(delta) + start
	return time.Unix(sec, 0)
}

// 根据生日获取年龄
func GetAgeByBirthDayYMD(birthDayYMD string) int {
	if len(birthDayYMD) != 8 {
		return 0
	}

	today := time.Now()
	birthDay, err := time.ParseInLocation("20060102", birthDayYMD, time.Local)
	if err != nil {
		return 0
	}

	age := today.Year() - birthDay.Year()
	if time.Date(today.Year(), birthDay.Month(), birthDay.Day(), 0, 0, 0, 0, time.Local).After(time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.Local)) {
		age -= 1
	}
	if age <= 0 {
		age = 1
	}

	return age
}
