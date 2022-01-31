package handler

import (
	"search-engine/pkg/filter"
	"search-engine/search"

	chttp "search-engine/pkg/occhttp"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UseCase search.UseCase
	Debug   bool
}

func NewHandler(useCase search.UseCase, debug bool) *Handler {
	return &Handler{
		UseCase: useCase,
		Debug:   debug,
	}
}

// GetSearch func it's a simple handler to search results
func (hdlr *Handler) GetSearch(ctx *fiber.Ctx) error {
	// fields alowed
	allowed := []string{"title", "text", "salary", "locality", "category"}

	qsURLString, err := filter.ParserQuery(ctx, allowed)
	if err != nil {
		ctx.Status(400)
		ctx.JSON(chttp.Failure(err.Error()))
		return nil
	}

	response, err := hdlr.UseCase.Find(qsURLString)
	if err != nil {
		return err
	}

	ctx.JSON(response)

	return nil
}
