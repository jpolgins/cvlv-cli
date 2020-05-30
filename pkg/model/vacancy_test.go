package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVacancy(t *testing.T) {
	vacancy := NewVacancy(
		"CTO",
		Salary{From: "1000", To: "2000"},
		"Google",
		"US",
		Period{PostedAt: "today", ExpiresAt: "tomorrow"},
	)

	assert.Equal(t, "CTO", vacancy.position)
	assert.Equal(t, "1000", vacancy.salary.From)
	assert.Equal(t, "2000", vacancy.salary.To)
	assert.Equal(t, "Google", vacancy.company)
	assert.Equal(t, "US", vacancy.location)
	assert.Equal(t, "today", vacancy.period.PostedAt)
	assert.Equal(t, "tomorrow", vacancy.period.ExpiresAt)
}

func TestVacancy_Company(t *testing.T) {
	vacancy := NewVacancy(
		"CTO",
		Salary{From: "1000", To: "2000"},
		"Goooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooogle",
		"US",
		Period{PostedAt: "today", ExpiresAt: "tomorrow"},
	)

	assert.Equal(t, "Gooooooooooooooooooooooooooooooooooooooooooooooooo", vacancy.Company())
}

func TestVacancy_Salary(t *testing.T) {
	vacancy := NewVacancy(
		"CTO",
		Salary{From: "1000", To: "2000"},
		"Google",
		"US",
		Period{PostedAt: "today", ExpiresAt: "tomorrow"},
	)

	assert.Equal(t, Salary{From: "1000", To: "2000"}, vacancy.Salary())
}

func TestVacancy_Position(t *testing.T) {
	vacancy := NewVacancy(
		"CTOooooooooooooooooooooooooooooooooooooooooooooooo",
		Salary{From: "1000", To: "2000"},
		"Google",
		"US",
		Period{PostedAt: "today", ExpiresAt: "tomorrow"},
	)

	assert.Equal(t, "CTOooooooooooooooooooooooooooooooooooooooooooooooo", vacancy.Position())
}

func TestVacancy_Period(t *testing.T) {
	vacancy := NewVacancy(
		"CTO",
		Salary{From: "1000", To: "2000"},
		"Google",
		"US",
		Period{PostedAt: "today", ExpiresAt: "tomorrow"},
	)

	assert.Equal(t, Period{PostedAt: "today", ExpiresAt: "tomorrow"}, vacancy.Period())
}

func TestSalary_String(t *testing.T) {
	assert.Equal(t, "1000 - 2000", fmt.Sprint(Salary{From: "1000", To: "2000"}))
	assert.Equal(t, "n/a - 2000", fmt.Sprint(Salary{From: "", To: "2000"}))
	assert.Equal(t, "1000 - n/a", fmt.Sprint(Salary{From: "1000", To: ""}))
	assert.Equal(t, "n/a", fmt.Sprint(Salary{From: "", To: ""}))
}

func TestPeriod_String(t *testing.T) {
	assert.Equal(t, "today - tomorrow", fmt.Sprint(Period{PostedAt: "today", ExpiresAt: "tomorrow"}))
	assert.Equal(t, "n/a - tomorrow", fmt.Sprint(Period{PostedAt: "", ExpiresAt: "tomorrow"}))
	assert.Equal(t, "today - n/a", fmt.Sprint(Period{PostedAt: "today", ExpiresAt: ""}))
	assert.Equal(t, "n/a", fmt.Sprint(Period{PostedAt: "", ExpiresAt: ""}))
}

func TestVacancy_MarshalJSON(t *testing.T) {
	vacancy := NewVacancy(
		"CTO",
		Salary{From: "1000", To: "2000"},
		"Google",
		"US",
		Period{PostedAt: "today", ExpiresAt: "tomorrow"},
	)

	vacancyJSON, _ := json.Marshal(vacancy)
	assert.Equal(t, string(vacancyJSON), `{"Position":"CTO","Salary":{"From":"1000","To":"2000"},"Company":"Google","Location":"US","Period":{"PostedAt":"today","ExpiresAt":"tomorrow"}}`)
}

func TestVacancy_UnmarshalJSON(t *testing.T) {
	jsn := []byte(`{"Position":"CTO","Salary":{"From":"1000","To":"2000"},"Company":"Google","Location":"US","Period":{"PostedAt":"today","ExpiresAt":"tomorrow"}}`)
	vacancyJSON := vacancyJSON{}
	if err := json.Unmarshal(jsn, &vacancyJSON); err != nil {
		panic(err)
	}

	assert.Equal(t, "CTO", vacancyJSON.Position)
	assert.Equal(t, "1000", vacancyJSON.Salary.From)
	assert.Equal(t, "2000", vacancyJSON.Salary.To)
	assert.Equal(t, "Google", vacancyJSON.Company)
	assert.Equal(t, "US", vacancyJSON.Location)
	assert.Equal(t, "today", vacancyJSON.Period.PostedAt)
	assert.Equal(t, "tomorrow", vacancyJSON.Period.ExpiresAt)
}
