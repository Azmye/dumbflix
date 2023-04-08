package handlers

import (
	episodeDto "dumbflix/dto/episode"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerEpisode struct {
	EpisodeRepository repositories.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerEpisode {
	return &handlerEpisode{EpisodeRepository}
}

func (h *handlerEpisode) FindEpisodes(c echo.Context) error {
	episodes, err := h.EpisodeRepository.FindEpisodes()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episodes})
}

func (h *handlerEpisode) GetEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}

func (h *handlerEpisode) CreateEpisode(c echo.Context) error {
	request := new(episodeDto.CreateEpisodeRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	episode := models.Episode{
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		VideoLink: request.VideoLink,
		MovieID:   request.MovieID,
	}

	data, err := h.EpisodeRepository.CreateEpisode(episode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func (h *handlerEpisode) UpdateMovie(c echo.Context) error {
	request := new(episodeDto.UpdateEpisodeRequest)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		episode.Title = request.Title
	}

	if request.Thumbnail != "" {
		episode.Thumbnail = request.Thumbnail
	}

	if request.VideoLink != "" {
		episode.VideoLink = request.VideoLink
	}

	if request.MovieID != 0 {
		if _, err := strconv.Atoi(strconv.Itoa(request.MovieID)); err == nil {
			episode.MovieID = request.MovieID
		}
	}

	data, err := h.EpisodeRepository.UpdateEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func (h *handlerEpisode) DeleteEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func convertResponseEpisode(u models.Episode) models.EpisodeResponse {
	return models.EpisodeResponse{
		ID:        u.ID,
		Title:     u.Title,
		Thumbnail: u.Thumbnail,
		VideoLink: u.VideoLink,
		Movie:     u.Movie,
	}
}
