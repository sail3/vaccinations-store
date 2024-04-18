package vaccination

import (
	"strings"
	"time"
)

type VaccinationRequest struct {
	ID     int        `json:"id,omitempty"`
	Name   string     `json:"name,omitempty"`
	DrugID int        `json:"drug_id,omitempty"`
	Dose   int        `json:"dose,omitempty"`
	Date   CustomTime `json:"date,omitempty"`
}

type CustomTime time.Time

func (c *CustomTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	*c = CustomTime(t)
	return nil
}

func (c CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}
