package route

import (
	"article/internal/delivery/http"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	App            *mux.Route
	PostController *http.PostController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/article", c.PostController.List)
	c.App.Post("/api/article", c.PostController.Create)
	c.App.Put("/api/article/:PostId", c.PostController.Update)
	c.App.Get("/api/article/:PostId", c.PostController.Get)
	c.App.Delete("/api/article/:PostId", c.PostController.Delete)
	c.App.Get("/api/articles", c.PostController.FindAllPost)

}
