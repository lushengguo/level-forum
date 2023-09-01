package fields

import (
	"fmt"
	"level-forum/fields/article"
)

type Fields struct {
	data map[string][]article.Article
}

func CreateField(name string, values []string) {
	// Add code to create a new field with the given name and values
	fmt.Println("calling CreateField")
}

func FindField(name string) []string {
	// Add code to find and return the values associated with the given field name
	return nil
}
