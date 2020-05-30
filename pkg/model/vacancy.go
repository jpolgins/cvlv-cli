package model

import (
	"encoding/json"
	"fmt"
)

const MaxContentLength = 50

type Vacancy struct {
	position string
	salary   Salary
	company  string
	location string
	period   Period
}

func NewVacancy(
	position string,
	salary Salary,
	company string,
	location string,
	period Period,
) *Vacancy {
	return &Vacancy{
		position: position,
		salary:   salary,
		company:  company,
		location: location,
		period:   period,
	}
}

func (v *Vacancy) Company() string {
	if len(v.company) > MaxContentLength {
		return v.company[:MaxContentLength]
	}

	return v.company
}

func (v *Vacancy) Salary() Salary {
	return v.salary
}

func (v *Vacancy) Position() string {
	if len(v.position) > MaxContentLength {
		return v.position[:MaxContentLength]
	}

	return v.position
}

func (v *Vacancy) Period() Period {
	return v.period
}

type Salary struct {
	From string
	To   string
}

func (s Salary) String() string {
	if s.From == "" && s.To == "" {
		return "n/a"
	}

	if s.From == "" {
		s.From = "n/a"
	}

	if s.To == "" {
		s.To = "n/a"
	}

	return fmt.Sprintf("%s - %s", s.From, s.To)
}

type Period struct {
	PostedAt  string
	ExpiresAt string
}

func (p Period) String() string {
	if p.PostedAt == "" && p.ExpiresAt == "" {
		return "n/a"
	}

	if p.PostedAt == "" {
		p.PostedAt = "n/a"
	}

	if p.ExpiresAt == "" {
		p.ExpiresAt = "n/a"
	}

	return fmt.Sprintf("%s - %s", p.PostedAt, p.ExpiresAt)
}

type vacancyJSON struct {
	Position string
	Salary   Salary
	Company  string
	Location string
	Period   Period
}

func (v *Vacancy) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(vacancyJSON{
		Position: v.position,
		Salary:   v.salary,
		Company:  v.company,
		Location: v.location,
		Period:   v.period,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (v *Vacancy) UnmarshalJSON(b []byte) error {
	j := vacancyJSON{
		Position: v.position,
		Salary:   v.salary,
		Company:  v.company,
		Location: v.location,
		Period:   v.period,
	}

	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	v.position = j.Position
	v.salary = j.Salary
	v.company = j.Company
	v.location = j.Location
	v.period = j.Period

	return nil
}