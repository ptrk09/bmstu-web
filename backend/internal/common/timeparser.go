package common

import (
	"fmt"
	"strconv"
	"time"
	"web/internal/model"
)

const (
	layout = "2006-01-02"
)

func GetCustomTime(year int, month int, day int) model.CustomTime {
	strMonth := strconv.Itoa(month)
	if month < 10 {
		strMonth = "0" + strMonth
	}

	strDay := strconv.Itoa(day)
	if day < 10 {
		strDay = "0" + strDay
	}

	stringDate := fmt.Sprintf(
		"%s-%s-%s",
		strconv.Itoa(year), strMonth, strDay,
	)
	fmt.Println(stringDate)
	updatedAt, error := time.Parse(layout, stringDate)
	if error != nil {
		fmt.Println(error)
	}

	return model.CustomTime{Time: updatedAt}
}
