package main

import (
	"fmt"
	"regexp"
	"time"
)

const layoutDate = "2006-01-02"
const (
	//MobileNumRegexp = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	PostCodeRegexp = "^[1-9][0-9]{5}$"
	//BirthdayRegexp  = "^((((1[6-9]|[2-9]\\d)\\d{2})-(0?[13578]|1[02])-(0?[1-9]|[12]\\d|3[01]))|(((1[6-9]|[2-9]\\d)\\d{2})-(0?[13456789]|1[012])-(0?[1-9]|[12]\\d|30))|(((1[6-9]|[2-9]\\d)\\d{2})-0?2-(0?[1-9]|1\\d|2[0-8]))|(((1[6-9]|[2-9]\\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00))-0?2-29-))$"
	BirthdayRegexp = "^[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}$"
)

var (
	MobileNumRegexp = regexp.MustCompile(`^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$`)
)

//验证手机号
func IsMobileNum(mobileNum string) bool {
	//mobileReg := regexp.MustCompile(MobileNumRegexp)
	//return mobileReg.MatchString(mobileNum)
	return MobileNumRegexp.MatchString(mobileNum)
}

//验证邮编
func IsPostCode(postCode string) bool {
	postCodeReg := regexp.MustCompile(PostCodeRegexp)
	return postCodeReg.MatchString(postCode)
}

//验证出生日期
func IsBirthday(birthday string) bool {
	birthdayReg := regexp.MustCompile(BirthdayRegexp)
	return birthdayReg.MatchString(birthday)
	//return len(s) >= 6 && len(s) <= 64 && EmailRegexp.MatchString(s)
}

//比较日期大小
func IsBeforeNow(birthday string) bool {
	birthdayDate, _ := time.Parse(layoutDate, birthday)
	nowDate := time.Now()
	return birthdayDate.Before(nowDate)
}

func main() {

	//fmt.Println(IsMobileNum("18867529162"))
	var a int8
	a = 127
	fmt.Println(a)
}
