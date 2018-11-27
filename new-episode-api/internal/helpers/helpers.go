package helpers

import (
	"fmt"
	"time"
)

func IsYesterday(date string) (bool, error) {
	const format = "2006-01-02"
	yesterday := time.Now().AddDate(0, 0, -1)

	formatted, err := time.Parse(format, date)
	if err != nil {
		return false, err
	}

	return yesterday.Format(format) == formatted.Format(format), nil
}

func ParseTvMazeUrl(id int) string {
	return fmt.Sprintf("http://api.tvmaze.com/shows/%d/episodes", id)
}
