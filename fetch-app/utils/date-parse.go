package utils

import (
	"strconv"
	"time"
)

func StringToTime(timeStr string) time.Time {

	sec, _ := strconv.ParseInt(timeStr, 10, 64)
	t := time.Unix(0, sec*int64(time.Millisecond))

	if year, _ := t.ISOWeek(); strconv.Itoa(year) == "0001" {
		layout := "2006-01-02T15:04:05.000Z"
		t, err := time.Parse(layout, timeStr)
		if err != nil {
			layout := "2006-01-02T15:04:05-07:00"
			t, err := time.Parse(layout, timeStr)
			if err != nil {
				layout := "2006-01-02T15:04:05-0700"
				t, err := time.Parse(layout, timeStr)
				if err != nil {
					layout := "2006-01-02T15:04:05-0700"
					t, _ := time.Parse(layout, timeStr)

					return t
				}

				return t
			}

			return t
		}

		return t

	} else {
		return t
	}

}
