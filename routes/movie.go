package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func MovieRoutes(e *echo.Group) {
	MovieRepository := repositories.RepositoryMovie(mysql.DB)

	h := handlers.HandlerMovie(MovieRepository)

	e.GET("/movies", h.FindMovies)
	e.GET("/movie/:id", h.GetMovie)
	e.POST("/movie", h.CreateMovie)
	e.PATCH("/movie/:id", h.UpdateMovie)
	e.DELETE("/movie/:id", h.DeleteMovie)
}
