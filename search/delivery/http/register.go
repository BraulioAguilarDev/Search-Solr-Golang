package handler

import (
	"search-engine/search"

	"github.com/gofiber/fiber/v2"
)

// MakeSearchHandler register our endpoints in a group router
func MakeSearchHandler(usecase search.UseCase, app *fiber.App, debug bool) {
	handlerSearch := NewHandler(usecase, debug)

	api := app.Group("/api/v1")
	api.Get("/search", handlerSearch.GetSearch)
}
