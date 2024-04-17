package vaccination

import "time"

type Vaccination struct {
	ID     int
	Name   string
	DrugID int
	Dose   int
	Date   time.Time
}
