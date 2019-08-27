package cronicleapi

import (
	"time"
)

// schedule timing

type TimingObject struct {
	Years    []int32    `json:"years,omitempty"`
	Months   []int32    `json:"months,omitempty"`
	Days     []int32    `json:"days,omitempty"`
	Weekdays []int32    `json:"weekdays,omitempty"`
	Hours    []int32    `json:"hours,omitempty"`
	Minutes  []int32    `json:"minutes,omitempty"`
}

// create minutely schedule 
func Minutely() *TimingObject {
	return &TimingObject {}
}

// create hourly schedule
func Hourly(minutes []int32) *TimingObject {
	return &TimingObject {
		Minutes: minutes,
	}
}

// create daily schedule
func Daily(minutes, hours []int32) *TimingObject {
	return &TimingObject {
		Minutes: minutes,
		Hours:   hours,
	}
}

// create weekly schedule
func Weekly(minutes, hours, weekdays []int32) *TimingObject {
	return &TimingObject {
		Weekdays: weekdays,
		Hours:    hours,
		Minutes:  minutes,
	}
}

// create monthly schedule
func Monthly(minutes, hours, days []int32) *TimingObject {
	return &TimingObject {
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
	}
}

// craete yearly schedule 
func Yearly(minutes, hours, days, months []int32) *TimingObject {
	return &TimingObject {
		Months:  months,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
	}
}

// create once schedule
func Oncely(t time.Time) *TimingObject {
	minute   := t.Minute()
	hour     := t.Hour()
	year,month,day := t.Date()
	return &TimingObject {
		Years:   []int32{ int32(year)   },
		Months:  []int32{ int32(month)  },
		Days:    []int32{ int32(day)    },
		Hours:   []int32{ int32(hour)   },
		Minutes: []int32{ int32(minute) },
	}
}
