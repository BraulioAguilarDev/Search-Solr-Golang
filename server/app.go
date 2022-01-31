package server

import (
	c "search-engine/config"
	handler "search-engine/search/delivery/http"
	shttp "search-engine/search/repository/solr"

	"search-engine/search/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang/glog"
)

// App struct
type App struct {
	DebugMode bool
	Service   *fiber.App

	SearchUseCase *usecase.SearchUseCase
}

// SetupRouter setup our endpoints
func (a *App) SetupRouter() {
	handler.MakeSearchHandler(a.SearchUseCase, a.Service, a.DebugMode)
}

// Run func executes our service
func (a *App) Run(port string) error {
	a.SetupRouter()

	if err := a.Service.Listen(":" + port); err != nil {
		glog.Error(err)

		return err
	}

	return nil
}

func NewApp(debug bool) (*App, error) {
	searchRepo := shttp.NewSearchRepository(c.Config.SearchHost, debug)

	fiber := fiber.New()
	fiber.Use(cors.New())

	app := &App{
		SearchUseCase: usecase.NewUseCase(searchRepo),
		DebugMode:     debug,
		Service:       fiber,
	}

	return app, nil
}
