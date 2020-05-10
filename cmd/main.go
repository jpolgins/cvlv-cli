package main

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/c-bata/go-prompt"
	"github.com/go-redis/redis/v7"
	"github.com/jpolgins/cvlv-cli/pkg/model"
	"github.com/jpolgins/cvlv-cli/pkg/source"
	"github.com/jpolgins/cvlv-cli/pkg/storage"
	"os"
)

var (
	vacancyRepository  model.VacancyRepository
	categoryRepository model.CategoryRepository
	categories         []model.Category
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	cvlv := source.NewCVLV()

	vacancyRepository = storage.NewVacancyRepository(redisClient, cvlv)
	categoryRepository = storage.NewCategoryRepository(redisClient, cvlv)
	categories = categoryRepository.GetAll()

	fmt.Println("Start typing")
	for {
		input := prompt.Input("> ", categoriesCompleter)

		if input == "exit" || input == "q" {
			os.Exit(0)
		}

		for _, category := range categories {
			if category.Title() == input {
				printTable(vacancyRepository.GetByCategory(category))
			}
		}
	}
}

func categoriesCompleter(d prompt.Document) []prompt.Suggest {
	suggestions := make([]prompt.Suggest, 0, 1)

	for _, category := range categories {
		suggestion := prompt.Suggest{Text: category.Title()}
		suggestions = append(suggestions, suggestion)
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func printTable(vacancies []model.Vacancy) {
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
