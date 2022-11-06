package model

const (
	layout = "2006-01-02"
)

// type CustomTime struct {
// 	time.Time
// }

// func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
// 	s := strings.Trim(string(b), "\"")
// 	date, err := time.Parse(layout, s)
// 	if err != nil {
// 		return err
// 	}
// 	t.Time = date

// 	return nil
// }

// func (j CustomTime) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(j.Time)
// }

// // Maybe a Format function for printing your date
// func (j CustomTime) Format(s string) string {
// 	t := j.Time
// 	return t.Format(s)
// }

type Player struct {
	ID      int        `json:"id"`
	FName   string     `json:"fname" binding:"required"`
	LNname  string     `json:"lname" binding:"required"`
	Country string     `json:"country" binding:"required"`
	Dob     CustomTime `json:"dob" binding:"required"`
}
