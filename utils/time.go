package utils

import(
	"time"
	"net/url"
)

func GetUnixTime(date string, format string) int64 {
	time, err := time.ParseInLocation(format, date, time.Local)
	if err != nil {
		return 0
	}

	return time.Unix()
}

func GetUrlTime(vars url.Values, parseDuration string) (int64, int64) {
	d, _ := time.ParseDuration(parseDuration)
	start := vars.Get("start")
	if start == "" {
		start = time.Now().Add(d).Format("2006-01-02 00:00")
	}

	end := vars.Get("end")
	if end == "" {
		end = time.Now().Format("2006-01-02") + " 23:59"
	}

	starttime := GetUnixTime(start, "2006-01-02 15:04")
	endtime := GetUnixTime(end, "2006-01-02 15:04") + 59

	vars.Set("start", start)
	vars.Set("end", end)

	return starttime, endtime
}
