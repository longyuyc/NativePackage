package public

import (
	"regexp"
	"strings"
	"time"
)

const (
	layoutDate      = "2006-01-02"
	EmailRegexp     = "^[0-9A-Za-z_.-]+@[0-9A-Za-z.-]+$"
	MobileNumRegexp = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	PostCodeRegexp  = "^[1-9][0-9]{5}$"
	DateRegexp      = "^((((1[6-9]|[2-9]\\d)\\d{2})-(0?[13578]|1[02])-(0?[1-9]|[12]\\d|3[01]))|(((1[6-9]|[2-9]\\d)\\d{2})-(0?[13456789]|1[012])-(0?[1-9]|[12]\\d|30))|(((1[6-9]|[2-9]\\d)\\d{2})-0?2-(0?[1-9]|1\\d|2[0-8]))|(((1[6-9]|[2-9]\\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00))-0?2-29-))$"
	DateTimeRegexp  = `^((((1[6-9]|[2-9]\d)\d{2})-(0?[13578]|1[02])-(0?[1-9]|[12]\d|3[01]))|(((1[6-9]|[2-9]\d)\d{2})-(0?[13456789]|1[012])-(0?[1-9]|[12]\d|30))|(((1[6-9]|[2-9]\d)\d{2})-0?2-(0?[1-9]|1\d|2[0-8]))|(((1[6-9]|[2-9]\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00))-0?2-29-)) (20|21|22|23|[0-1]?\d):[0-5]?\d:[0-5]?\d$`
)

//校验手机号
func IsMobileNum(mobileNum string) bool {
	mobileReg := regexp.MustCompile(MobileNumRegexp)
	return mobileReg.MatchString(mobileNum)
}

//校验邮编
func IsPostCode(postCode string) bool {
	postCodeReg := regexp.MustCompile(PostCodeRegexp)
	return postCodeReg.MatchString(postCode)
}

//校验邮箱
func IsEmail(email string) bool {
	emailReg := regexp.MustCompile(EmailRegexp)
	return len(email) >= 6 && len(email) <= 64 && emailReg.MatchString(email)
}

//校验出生日期
func IsBirthday(birthday string) bool {
	birthdayReg := regexp.MustCompile(DateRegexp)
	return IsBeforeNow(birthday) && birthdayReg.MatchString(birthday)
}

//校验收货人地址信息
func IsDetail(detail string) bool {
	if len(detail) == 0 || len(detail) > 255 {
		return false
	} else {
		return true
	}
}

//校验收件人姓名
func IsName(name string) bool {
	if len(name) == 0 || len(name) > 255 {
		return false
	} else {
		return true
	}
}

//校验日期
func IsValidDate(date string) bool {
	dateReg := regexp.MustCompile(DateRegexp)
	return dateReg.MatchString(date)
}

//校验日期+时间
func IsDateTime(dateTime string) bool {
	if strings.Index(dateTime, "/") > 0 {
		dateTime = strings.Replace(dateTime, "/", "-", 2)
	}
	dateTimeReg := regexp.MustCompile(DateTimeRegexp)
	return dateTimeReg.MatchString(dateTime)
}

//比较输入日期与当前日期大小
func IsBeforeNow(birthday string) bool {
	birthdayDate, _ := time.Parse(layoutDate, birthday)
	nowDate := time.Now()
	return birthdayDate.Before(nowDate)
}
