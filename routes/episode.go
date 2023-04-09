package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoutes(e *echo.Group) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)

	h := handlers.HandlerEpisode(EpisodeRepository)

	e.GET("/movie/:movieID/episodes", middleware.Auth(h.FindEpisodes))
	e.GET("/movie/:movieID/episode/:id", middleware.Auth(h.GetEpisode))
	e.POST("/episode", middleware.Auth(h.CreateEpisode))
	e.PATCH("/episode/:id", middleware.Auth(h.UpdateEpisode))
	e.DELETE("/episode/:id", middleware.Auth(h.DeleteEpisode))
}
