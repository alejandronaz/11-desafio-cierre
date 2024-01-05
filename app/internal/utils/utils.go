package utils

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrTimeIncorrect = errors.New("time provided is incorrect")
)

func ConvertTime(time string) (hour, min int, err error) {

	// get hour and min
	hourStr, minStr, found := strings.Cut(time, ":")
	if !found {
		return 0, 0, ErrTimeIncorrect
	}

	// parse to int both
	hour, err = strconv.Atoi(hourStr)
	if err != nil {
		return 0, 0, ErrTimeIncorrect
	}
	min, err = strconv.Atoi(minStr)
	if err != nil {
		return 0, 0, ErrTimeIncorrect
	}

	// check time
	if hour < 0 || hour > 23 || min < 0 || min > 59 {
		return 0, 0, ErrTimeIncorrect
	}

	return hour, min, nil
}
