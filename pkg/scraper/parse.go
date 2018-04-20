package scraper

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseDate(date string) (startTime, endTime time.Time, err error) {
	// MONTH DD-MONTH DD, YYYY
	matches := regexp.MustCompile(`(\w+)\s(\d+).(\w+)\s(\d+)\,\s(\d+)`).FindStringSubmatch(date)

	if len(matches) != 6 {
		// MONTH DD-DD, YYYY
		matches = regexp.MustCompile(`(\w+)\s(\d+).(\d+)\,\s(\d+)`).FindStringSubmatch(date)
		if len(matches) != 5 {
			// MONTH DD, YYYY
			matches = regexp.MustCompile(`(\w+)\s(\d+).*\,\s(\d+)`).FindStringSubmatch(date)
			if len(matches) != 4 {
				return time.Time{}, time.Time{}, errors.New(`expected format: "May 1-2, 2018" or "May 1 2018" from: ` + date)
			}
		}
	}

	startDay, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endDay, err := strconv.Atoi(matches[len(matches)-2])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startMonth, err := parseMonth(matches[1])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	var endMonth time.Month
	// MONTH DD-MONTH DD, YYYY
	if len(matches) == 6 {
		endMonth, err = parseMonth(matches[3])
		if err != nil {
			return time.Time{}, time.Time{}, err
		}

		// MONTH DD-DD, YYYY || MONTH DD, YYYY
	} else {
		endMonth = startMonth
	}

	year, err := strconv.Atoi(matches[len(matches)-1])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startTime = time.Date(year, startMonth, startDay, 0, 0, 0, 0, time.UTC)
	endTime = time.Date(year, endMonth, endDay, 0, 0, 0, 0, time.UTC)

	return startTime, endTime, nil
}

func parseMonth(month string) (time.Month, error) {
	for i := 1; i <= 12; i++ {
		if strings.EqualFold(month, time.Month(i).String()) {
			return time.Month(i), nil
		}
	}
	return 0, errors.New("invalid month: " + month)
}

func parseLocation(location string) (country, state, city string, err error) {
	if len(strings.Split(location, ", ")) != 2 {
		return "", "", "", fmt.Errorf("%s does not contain `, `", location)
	}
	country = "United States"
	state = strings.Split(location, ", ")[1]
	city = strings.Split(location, ", ")[0]
	return country, state, city, err
}
