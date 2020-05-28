package scrapper

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/jpolgins/cvlv-cli/pkg/cache"
	"github.com/jpolgins/cvlv-cli/pkg/model"
	"regexp"
	"strings"
)

type Scrapper struct {
	crawler *colly.Collector
	cache   cache.Cache
}

func NewScrapper(cache cache.Cache) *Scrapper {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.cv.lv"),
		colly.AllowURLRevisit(),
	)

	collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nVisiting ====> %s ", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("[Done]")
	})

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("error:", e, r.Request.URL, string(r.Body))
	})

	return &Scrapper{
		crawler: collector,
		cache:   cache,
	}
}

func (s *Scrapper) FetchVacanciesBy(category model.Category) []model.Vacancy {
	var vacancies []model.Vacancy
	fromCache, has := s.cache.Get(category.URI())

	if !has {
		s.crawler.OnHTML(".cvo_module_offers_wrap .cvo_module_offer", func(e *colly.HTMLElement) {
			position := e.DOM.Find(".offer_primary_info h2 a").Text()
			salary := parseSalary(e.DOM.Find(".offer-salary").Text())
			company := e.DOM.Find(".offer-company").Text()
			location := e.DOM.Find(".offer-location").Text()
			period := model.Period{
				PostedAt:  e.DOM.Find(".offer_dates li").First().Text(),
				ExpiresAt: e.DOM.Find(".offer_dates li").Next().Text(),
			}

			vacancy := model.NewVacancy(position, salary, company, location, period)
			vacancies = append(vacancies, *vacancy)
		})

		s.crawler.OnHTML(".page_next a[href]", func(e *colly.HTMLElement) {
			if err := e.Request.Visit(e.Attr("href")); err != nil {
				panic(err)
			}
		})

		if err := s.crawler.Visit(fmt.Sprintf("https://www.cv.lv/job-ads/%s", category.URI())); err != nil {
			panic(err)
		}

		serialized, _ := json.Marshal(vacancies)
		s.cache.Set(category.URI(), serialized, 0)

		return vacancies
	}

	if err := json.Unmarshal([]byte(fmt.Sprintf("%v", fromCache)), &vacancies); err != nil {
		panic(err) // TODO: return error
	}

	return vacancies
}

func (s *Scrapper) FetchCategories() []model.Category {
	var categories []model.Category
	fromCache, has := s.cache.Get("categories")

	if !has {
		s.crawler.OnHTML("#select-field option", func(e *colly.HTMLElement) {
			title := e.DOM.Text()
			URI := e.Attr("value")
			categories = append(categories, *model.NewCategory(title, URI))
		})

		if err := s.crawler.Visit("https://www.cv.lv/english"); err != nil {
			panic(err)
		}

		serialized, _ := json.Marshal(categories)
		s.cache.Set("categories", serialized, 0)

		return categories
	}

	if err := json.Unmarshal([]byte(fmt.Sprintf("%v", fromCache)), &categories); err != nil {
		panic(err) // TODO: return error
	}

	return categories
}

func parseSalary(str string) model.Salary {
	salary := model.Salary{}
	re := regexp.MustCompile(`(?m)(?:[^\d,]|^)(\d+(?:(?:,\d+)*,\d{3})?\.\d{2,3})\b`)

	for i, match := range re.FindAllString(str, -1) {
		if len(match) != 0 && i == 0 {
			salary.From = strings.Trim(match, " ")
		}

		if len(match) != 0 && i == 1 {
			salary.To = strings.Trim(match, " ")
		}
	}

	return salary
}
