package util

import (
	"strconv"
	"strings"
	"time"

	"d2din.com/internal/pkg/log"
)

func ChangeToUTC(lasting time.Time) time.Time {
	return lasting.UTC()
}

func ChangeToRFC3339(lasting time.Time) string {
	return lasting.Format(time.RFC3339)
}

func NowToUTC() time.Time {
	return time.Now().UTC()
}

func UTCToRFC3339() string {
	utc := NowToUTC()
	return ChangeToRFC3339(utc)
}

func GetAge(birthday time.Time) (age int) {
	if birthday.IsZero() {
		return 0
	}

	now := NowToUTC()
	age = now.Year() - birthday.Year()
	if int(now.Month()) < int(birthday.Month()) || int(now.Day()) < int(birthday.Day()) {
		age--
	}
	return age
}

func BydStringToUTC(strTime string) (tm time.Time) {
	strTime = strings.ReplaceAll(strTime, "/Date(", "")
	strTime = strings.ReplaceAll(strTime, ")/", "")
	temp, err := strconv.ParseInt(strTime, 10, 64)
	if err != nil {
		log.Error(err)
	}

	tm = time.Unix(temp/1000, 0)
	if err != nil {
		log.Error(err)
	}
	return tm
}

func ChangUnixTimeStringToUTC(unixTimeString string) (tm time.Time) {
	i, err := strconv.ParseInt(unixTimeString, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(i, 0)
}
