package search

import (
	"github.com/a-h/templ"
)

func InitSearch() *templ.ComponentHandler {
	component := main("John")

	return templ.Handler(component)
}
