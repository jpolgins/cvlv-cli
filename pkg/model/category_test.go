package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCategory(t *testing.T) {
	category := NewCategory("SomeTitle", "https://some.uri")

	assert.Equal(t, "SomeTitle", category.title)
	assert.Equal(t, "https://some.uri", category.uri)
}

func TestCategory_Title(t *testing.T) {
	category := NewCategory("SomeTitle", "https://some.uri")

	assert.Equal(t, "SomeTitle", category.title)
}

func TestCategory_URI(t *testing.T) {
	category := NewCategory("SomeTitle", "https://some.uri")

	assert.Equal(t, "https://some.uri", category.uri)
}

func TestCategory_MarshalJSON(t *testing.T) {
	category := NewCategory("SomeTitle", "https://some.uri")
	categoryJSON, _ := json.Marshal(category)

	assert.Equal(t, string(categoryJSON), `{"Title":"SomeTitle","URI":"https://some.uri"}`)
}

func TestCategory_UnmarshalJSON(t *testing.T) {
	jsn := []byte(`{"Title":"SomeTitle","URI":"https://some.uri"}`)
	categoryJSON := categoryJSON{}

	if err := json.Unmarshal(jsn, &categoryJSON); err != nil {
		panic(err)
	}

	assert.Equal(t, "SomeTitle", categoryJSON.Title)
	assert.Equal(t, "https://some.uri", categoryJSON.URI)
}
