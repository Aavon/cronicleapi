package cronicleapi

import (
	"strings"
	"time"
	"strconv"
	"log"
)

// 重复类型
type RepeatType int
const (
	Once RepeatType = iota
	Day
	Month
	Year
	Week
	WorkDay
)

// 解析规则，生成cron表达式

func GenCronExp(noticeRule, repeatType string, startTime time.Time) string {
	exp := ""
	na := strings.Split(noticeRule,"^")
	if len(na) == 2 && na[0] != "N" {
		// 解析提醒时间
		switch (na[0]) {
		case "F":
			// 计算提前后的时间
			mins,err := strconv.ParseInt(na[1],10,64)
			if err != nil {
				log.Println(err)
				return exp
			}
			mins = -1 * mins
			startTime = startTime.Add(time.Duration(mins) * time.Minute)
		case "T":
			// 具体提醒时间
			stamp,err := strconv.ParseInt(na[1],10,64)
			if err != nil {
				log.Println(err)
				return exp
			}
			startTime = time.Unix(stamp,0)
		}

		flags := []string{}
		// 秒
		flags = append(flags,strconv.FormatInt(int64(startTime.Second()),10))
		// 分
		flags = append(flags,strconv.FormatInt(int64(startTime.Minute()),10))
		// 时
		flags = append(flags,strconv.FormatInt(int64(startTime.Hour()),10))
		y,m,d := startTime.Date()
		// 日,月
		flags = append(flags,strconv.FormatInt(int64(d),10))
		flags = append(flags,strconv.FormatInt(int64(m),10))
		// 周
		wf := int(startTime.Weekday()) 
		// sunday -- 0
		wf = wf + 1
		flags = append(flags,strconv.FormatInt(int64(wf),10))
		// 年
		flags = append(flags,strconv.FormatInt(int64(y),10))
		// 解析重复次数
		switch (repeatType) {
		case "once":
			// do nothing
			flags[5] = "?"
		case "day":
			flags[3] = "*"
			flags[4] = "*"
			flags[5] = "?"
			flags[6] = ""
		case "month":
			flags[4] = "*"
			flags[5] = "?"
			flags[6] = ""
		case "year":
			flags[5] = "?"
			flags[6] = ""
		case "week":
			flags[3] = "?"
			flags[4] = "*"
			flags[6] = ""
		case "work_day":
			flags[3] = "?"
			flags[4] = "*"
			flags[5] = "2-6"
			flags[6] = ""
		}
		exp = strings.Join(flags," ")
	}
	return exp;
}


// 解析规则,生成Toming对象
func GenTimming(repeatType RepeatType, startTime time.Time) *TimingObject {

	flags := []int32{}
	// 秒
	flags = append(flags, int32(startTime.Second()))
	// 分
	flags = append(flags, int32(startTime.Minute()))
	// 时
	flags = append(flags, int32(startTime.Hour()))
	y,m,d := startTime.Date()
	// 日,月
	flags = append(flags, int32(d))
	flags = append(flags, int32(m))
	// 周
	wf := int(startTime.Weekday()) 
	// sunday -- 0
	flags = append(flags, int32(wf))
	// 年
	flags = append(flags, int32(y))
	var timing *TimingObject
	switch (repeatType) {
	case Once:
		timing = Oncely(startTime)
	case Day:
		timing = Daily([]int32{flags[1]}, []int32{flags[2]})
	case Month:
		timing = Monthly([]int32{flags[1]}, []int32{flags[2]}, []int32{flags[3]})
	case Year:
		timing = Yearly([]int32{flags[1]}, []int32{flags[2]}, []int32{flags[3]}, []int32{flags[4]})
	case Week:
		timing = Weekly([]int32{flags[1]}, []int32{flags[2]}, []int32{flags[5]})
	case WorkDay:
		timing = Weekly([]int32{flags[1]}, []int32{flags[2]}, []int32{1,2,3,4,5})
	}
	return timing
}

