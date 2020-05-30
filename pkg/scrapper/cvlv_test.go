package scrapper

import (
	"github.com/jpolgins/cvlv-cli/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScrapper(t *testing.T) {
	s := NewScrapper(Options{})
	assert.NotNil(t, s)
}

func TestScrapper_FetchCategories(t *testing.T) {
	s := NewScrapper(Options{})
	res, err := s.FetchCategories()

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestScrapper_FetchVacanciesBy(t *testing.T) {
	s := NewScrapper(Options{})
	category, _ := s.FetchCategories()
	vacancy, err := s.FetchVacanciesBy(category[0])

	assert.NoError(t, err)
	assert.NotNil(t, vacancy)
}

func Test_parseSalary(t *testing.T) {
	s := parseSalary("Monthly salary: 800.00 to 6000.00 EUR")
	assert.Equal(t, model.Salary{
		From: "800.00",
		To:   "6000.00",
	}, s)
}