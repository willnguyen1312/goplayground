package books

import (
	"encoding/json"
	"errors"
	"strings"
)

// Book struct
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// CategoryByLength method on Book struct
func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

// AuthorLastName method on Book struct
func (b *Book) AuthorLastName() string {
	return strings.Fields(b.Author)[1]
}

// NewBookFromJSON function
func NewBookFromJSON(jsonInput string) (Book, error) {
	result := Book{}

	if err := json.Unmarshal([]byte(jsonInput), &result); err != nil {
		return result, errors.New("Oops")
	}
	return result, nil
}
