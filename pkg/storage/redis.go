package storage

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/jpolgins/cvlv-cli/pkg/model"
	"github.com/jpolgins/cvlv-cli/pkg/source"
)

type categoryRepository struct {
	redis *redis.Client
	cvlv  *source.CVLV
}

func NewCategoryRepository(redis *redis.Client, cvlv *source.CVLV) model.CategoryRepository {
	return &categoryRepository{redis, cvlv}
}

func (c *categoryRepository) GetAll() []model.Category {
	fromRedis, _ := c.redis.Get("categories").Result()

	if len(fromRedis) == 0 {
		categories := c.cvlv.FetchCategories()
		serialized, _ := json.Marshal(categories)
		c.redis.Set("categories", serialized, 0)

		return categories
	}

	var categories []model.Category
	if err := json.Unmarshal([]byte(fromRedis), &categories); err != nil {
		panic(err)
	}

	return categories
}

type vacancyRepository struct {
	redis *redis.Client
	cvlv  *source.CVLV
}

func NewVacancyRepository(redis *redis.Client, cvlv *source.CVLV) model.VacancyRepository {
	return &vacancyRepository{redis, cvlv}
}

func (v *vacancyRepository) GetByCategory(category model.Category) []model.Vacancy {
	fromRedis, _ := v.redis.Get(category.URI()).Result()

	if len(fromRedis) == 0 {
		vacancies := v.cvlv.FetchVacanciesBy(category)
		serialized, _ := json.Marshal(vacancies)
		v.redis.Set(category.URI(), serialized, 0)

		return vacancies
	}

	var vacancies []model.Vacancy
	_ = json.Unmarshal([]byte(fromRedis), &vacancies)

	return vacancies
}
