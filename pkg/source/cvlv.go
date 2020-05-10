package source

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/jpolgins/cvlv-cli/pkg/model"
	"regexp"
	"strings"
)

type CVLV struct {
	crawler *colly.Collector
}

func NewCVLV() *CVLV {
	c := colly.NewCollector(
		colly.AllowedDomains("www.cv.lv"),
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nVisiting ====> %s ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[Done]")
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("error:", e, r.Request.URL, string(r.Body))
	})

	return &CVLV{
		crawler: c,
	}
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

func (c *CVLV) FetchVacanciesBy(category model.Category) []model.Vacancy {
	var vacancies []model.Vacancy

	c.crawler.OnHTML(".cvo_module_offers_wrap .cvo_module_offer", func(e *colly.HTMLElement) {
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

	c.crawler.OnHTML(".page_next a[href]", func(e *colly.HTMLElement) {
		if err := e.Request.Visit(e.Attr("href")); err != nil {
			panic(err)
		}
	})

	categoryURL := fmt.Sprintf("https://www.cv.lv/job-ads/%s", category.URI())
	if err := c.crawler.Visit(categoryURL); err != nil {
		panic(err)
	}

	return vacancies
}

func (c *CVLV) FetchCategories() []model.Category {
	var categories []model.Category
	c.crawler.OnHTML("#select-field option", func(e *colly.HTMLElement) {
		title := e.DOM.Text()
		URI := e.Attr("value")
		category := model.NewCategory(title, URI)
		categories = append(categories, *category)
	})

	err := c.crawler.Visit("https://www.cv.lv/english")

	if err != nil {
		panic(err)
	}

	return categories
}
