package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoutes(e *echo.Group) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)

	h := handlers.HandlerEpisode(EpisodeRepository)

	e.GET("/movie/:movieID/episodes", h.FindEpisodes)
	e.GET("/movie/:movieID/episode/:id", h.GetEpisode)
	e.POST("/episode", h.CreateEpisode)
	e.PATCH("/episode/:id", h.UpdateEpisode)
	e.DELETE("/episode/:id", h.DeleteEpisode)
}
