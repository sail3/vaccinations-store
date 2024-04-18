package drug

import (
	"strings"
	"time"
)

type RegisterDrugRequest struct {
	Name        string     `json:"name,omitempty"`
	Approved    bool       `json:"approved,omitempty"`
	MinDose     int        `json:"min_dose,omitempty"`
	MaxDose     int        `json:"max_dose,omitempty"`
	AvailableAt CustomTime `json:"available_at,omitempty"`
}

type UpdateDrugRequest struct {
	Name        string     `json:"name"`
	Approved    bool       `json:"approved"`
	MinDose     int        `json:"minDose"`
	MaxDose     int        `json:"maxDose"`
	AvailableAt CustomTime `json:"availableAt"`
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
