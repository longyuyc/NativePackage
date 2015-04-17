package LyTime

import (
	"fmt"
	"time"
)

const layoutTime = "2006-01-02 15:04:05"

func FormatUseSecond() {
	t := time.Unix(1362984425, 0)
	nt := t.Format(layoutTime)
	fmt.Println(nt)
}

func GetNowTime() {
	//获取当前时间的结构体
	t := time.Now()
	f := t.Format(layoutTime)
	fmt.Println(f)
}

//获取当前时间的时间戳
func GetTimeStamp() {
	//获取当前时间时间戳
	t := time.Now()
	fmt.Println(t.Unix())
	//获取指定时间时间戳
	appoint := time.Date(2015, 4, 9, 17, 27, 20, 0, time.Local)
	fmt.Println(appoint.Unix())
}

//时间比较
func CompareTime() {
	t := time.Date(2015, 1, 2, 0, 0, 0, 0, time.Local)
	b := time.Date(2000, 1, 1, 8, 0, 0, 0, time.Local)
	//如果t的时间在b之前返回ture,否则返回false
	fmt.Println(t.Before(b))
	//如果t的时间在b的时间之后，返回ture,否则返回false
	fmt.Println(t.After(b))
}

//UTC对应0时区
//LoadLocation返回使用给定的名字创建的Location。
func GetLoadLocation() {
	loc := new(time.Location)
	loc, _ = time.LoadLocation("UTC")
	fmt.Println(loc)
}

//FixedZone使用给定的地点名name和时间偏移量offset（单位秒）创建并返回一个Location
func GetFixedZone() {
	loc := time.FixedZone("UTC", 1)
	fmt.Println(loc)
}

//Unix创建一个本地时间，对应sec和nsec表示的Unix时间（从January 1, 1970 UTC至该时间的秒数和纳秒数）。
//nsec的值在[0, 999999999]范围外是合法的。
func GetUnix() {
	fmt.Println(time.Now().Unix())
}

//Location返回t的地点和时区信息。
func GetLocation() {
	t := time.Now()
	loc := t.Location()
	fmt.Println(loc)
}

//Zone计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）。
func GetZone() {
	t := time.Now()
	name, offset := t.Zone()
	fmt.Println(name)
	fmt.Println(offset)
}

//IsZero报告t是否代表Time零值的时间点，January 1, year 1, 00:00:00 UTC。
func GetIsZero() {
	t := time.Now()
	fmt.Println(t.IsZero())
}

//UTC返回采用UTC和零时区，但指向同一时间点的Time。
func GetUTC() {
	t := time.Now()
	u := t.UTC()
	fmt.Println(u)
}

/*
func GetIn() {
	t := time.Now()
	loc := &time.Location
	in := t.In(loc)
}

 UnixNano
*/

func GetUnixNano() {
	t := time.Now()
	nano := t.UnixNano()
	fmt.Println(nano)
}

//判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。
//本方法和用t==u不同，这种方法还会比较地点和时区信息。

func GetEqual() {
	t := time.Date(2015, 4, 9, 10, 0, 0, 0, time.Local)
	u := time.Date(2015, 4, 9, 10, 0, 0, 0, time.Local)
	r := t.Equal(u)
	fmt.Print(r)
}

//返回t对应的那一天的时、分、秒。
func GetClock() {
	t := time.Now()
	hour, min, sec := t.Clock()
	fmt.Println(hour)
	fmt.Println(min)
	fmt.Println(sec)
}

//返回时间点t对应的年份
func GetYear() {
	t := time.Now()
	y := t.Year()
	fmt.Println(y)
}

//返回时间点t对应那一年的第几月
func GetMonth() {
	t := time.Now()
	m := t.Month()
	fmt.Println(m)
}

//返回时间点t对应的ISO 9601标准下的年份和星期编号。
//星期编号范围[1,53]，1月1号到1月3号可能属于上一年的最后一周，
//12月29号到12月31号可能属于下一年的第一周。
func GetISOWeek() {
	t := time.Now()
	year, isoWeek := t.ISOWeek()
	fmt.Println(year)
	fmt.Println(isoWeek)
}

//返回时间点t对应的那一年的第几天，平年的返回值范围[1,365]，闰年[1,366]。
func GetYearDay() {
	t := time.Now()
	yearDay := t.YearDay()
	fmt.Println(yearDay)
}

//返回时间点t对应那一月的第几日。
func GetDay() {
	t := time.Now()
	day := t.Day()
	fmt.Println(day)
}

//返回时间点t对应的那一周的周几。
func GetWeekDay() {
	t := time.Now()
	weekDay := t.Weekday()
	fmt.Println(weekDay)
}

//返回t对应的那一天的第几小时，范围[0, 23]。
func GetHour() {
	t := time.Now()
	hour := t.Hour()
	fmt.Println(hour)
}

//返回t对应的那一小时的第几分种，范围[0, 59]。
func GetMinute() {
	t := time.Now()
	minute := t.Minute()
	fmt.Println(minute)
}

//返回t对应的那一分钟的第几秒，范围[0, 59]。
func GetSecond() {
	t := time.Now()
	second := t.Second()
	fmt.Println(second)
}

//返回t对应的那一秒内的纳秒偏移量，范围[0, 999999999]。
func GetNanosecond() {
	t := time.Now()
	nanosecond := t.Nanosecond()
	fmt.Println(nanosecond)
}

//Add返回时间点t+d
func GetAdd() {
	t := time.Now()
	d := t.Add(1)
	fmt.Println(d)
}

/*
AddDate返回增加了给出的年份、月份和天数的时间点Time。例如，时间点January 1, 2011调用AddDate(-1, 2, 3)会返回March 4, 2010。
AddDate会将结果规范化，类似Date函数的做法。因此，举个例子，给时间点October 31添加一个月，会生成时间点December 1。（从时间点November 31规范化而来）
*/
func GetAddDate() {
	t := time.Now()
	ad := t.AddDate(-1, 0, 0)
	fmt.Println(ad)
}

/*
返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，
将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。
*/
func GetSub() {
	t := time.Now()
	s := time.Now()
	r := t.Sub(s.AddDate(0, 0, -1))
	fmt.Println(r)
}

/*
返回距离t最近的时间点，该时间点应该满足从Time零值到该时间点的时间段能整除d；如果有两个满足要求的时间点，距离t相同，会向上舍入；如果d <= 0，会返回t的拷贝。
*/
func GetRound() {
	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}
	for _, d := range round {
		fmt.Printf("t.Round(%6s)=%s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
}

/*
Format根据layout指定的格式返回t代表的时间点的格式化文本表示。layout定义了参考时间：
格式化后的字符串表示，它作为期望输出的例子。同样的格式规则会被用于格式化时间。
预定义的ANSIC、UnixDate、RFC3339和其他版式描述了参考时间的标准或便捷表示。
要获得更多参考时间的定义和格式，参见本包的ANSIC和其他版式常量。
*/
func GetFormat() {
	t := time.Now()
	fmt.Println(t.Format(layoutTime))
	fmt.Println(t.UTC().Format(layoutTime))
}

//String返回采用如下格式字符串的格式化时间。
func GetString() {
	t := time.Now()
	fmt.Println(t.String())
}

//Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回。
func GetSleep() {
	s := time.Now()
	fmt.Println(s)
	time.Sleep(1000 * time.Millisecond)
	e := time.Now()
	fmt.Println(e)
}
