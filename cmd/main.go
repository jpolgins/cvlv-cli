package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/go-redis/redis/v7"
	"github.com/jpolgins/cvlv-cli/pkg/cache"
	"github.com/jpolgins/cvlv-cli/pkg/scrapper"
	"github.com/jpolgins/cvlv-cli/pkg/ui"
	"github.com/umputun/go-flags"
	"os"
)

func main() {
	var opts Opts
	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     opts.Redis.Addr,
		Password: opts.Redis.Passwd,
		DB:       opts.Redis.DB,
	})

	cvlv := scrapper.NewScrapper(cache.NewRedisCache(redisClient))
	categories := cvlv.FetchCategories()

	// Add suggestions to CLI completer.
	completer := func(d prompt.Document) []prompt.Suggest {
		suggestions := make([]prompt.Suggest, 0, 1)

		for _, category := range categories {
			suggestion := prompt.Suggest{Text: category.Title()}
			suggestions = append(suggestions, suggestion)
		}

		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}

	for {
		input := prompt.Input("> ", completer)

		if input == "exit" || input == "q" || input == "q!" {
			os.Exit(0)
		}

		for _, category := range categories {
			if category.Title() == input {
				ui.Print(cvlv.FetchVacanciesBy(category))
			}
		}
	}
}