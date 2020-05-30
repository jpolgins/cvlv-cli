package model

import "encoding/json"

type Category struct {
	title string
	uri   string
}

func NewCategory(title string, uri string) *Category {
	return &Category{
		title: title,
		uri:   uri,
	}
}

func (c *Category) Title() string {
	return c.title
}

func (c *Category) URI() string {
	return c.uri
}

type categoryJSON struct {
	Title string
	URI   string
}

func (c *Category) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(categoryJSON{
		Title: c.title,
		URI:   c.uri,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *Category) UnmarshalJSON(b []byte) error {
	j := categoryJSON{
		Title: c.title,
		URI:   c.uri,
	}

	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	c.title = j.Title
	c.uri = j.URI

	return nil
}
