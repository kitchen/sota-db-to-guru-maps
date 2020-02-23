package sotadb

import "time"

const (
	ISO8601DateFormat = "2006-01-02"
	DBDateFormat      = "02/01/2006"
)

type SummitDate struct {
	time.Time
}

func (date *SummitDate) UnmarshalCSV(csv string) (err error) {
	if csv == "" {
		return nil
	}

	date.Time, err = time.Parse(DBDateFormat, csv)
	return err
}

func (date *SummitDate) MarshalCSV() (string, error) {
	if date.IsZero() {
		return "", nil
	}
	return date.Time.Format(DBDateFormat), nil
}
