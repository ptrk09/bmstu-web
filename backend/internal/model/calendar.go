package model

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	t.Time = date

	return nil
}

func (j CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Time)
}

// Maybe a Format function for printing your date
func (j CustomTime) Format(s string) string {
	t := j.Time
	return t.Format(s)
}

type Calendar struct {
	ID        int        `json:"id" db:"id"`
	ListingID int        `json:"listingID" db:"listing_id" binding:"required"`
	Date      CustomTime `json:"date" binding:"required"`
	Available bool       `json:"available" binding:"required"`
}
