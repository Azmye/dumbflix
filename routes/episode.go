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

	e.GET("/movie/:movieID/episodes", middleware.Auth(h.FindEpisodesByMovie))
	e.GET("/movie/:movieID/episode/:id", middleware.Auth(h.GetEpisodeByMovie))
	e.POST("/episode", middleware.Auth(middleware.UploadFile(h.CreateEpisode)))
	e.PATCH("/episode/:id", middleware.Auth(middleware.UploadFile(h.UpdateEpisode)))
	e.DELETE("/episode/:id", middleware.Auth(h.DeleteEpisode))
}
