package handlers

import (
	dto "dumbflix/dto/result"
	"dumbflix/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerMovie struct {
	MovieRepository repositories.MovieRepository
}

func (h *handlerMovie) FindMovies(c echo.Context) error {
	movies, err := h.MovieRepository.FindMovies()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: movies})
}
