package ui

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/jpolgins/cvlv-cli/pkg/model"
)

func Print(vacancies []model.Vacancy) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "POSITION"},
			{Align: simpletable.AlignCenter, Text: "COMPANY"},
			{Align: simpletable.AlignCenter, Text: "POSTED/EXPIRES"},
			{Align: simpletable.AlignCenter, Text: "SALARY €"},
		},
	}

	for k, v := range vacancies {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", k+1)},
			{Text: v.Position()},
			{Text: v.Company()},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(v.Period())},
			{Align: simpletable.AlignLeft, Text: fmt.Sprint(v.Salary())},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table.String())
}
