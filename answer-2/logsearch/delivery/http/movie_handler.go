package http

import (
	"strconv"

	"github.com/khihadysucahyo/golang-test/answer-2/domain"
	"github.com/khihadysucahyo/golang-test/answer-2/movie/usecase"
	"github.com/labstack/echo/v4"
)

type movieHandler struct {
	movieUsecase usecase.MovieUsecase
}

// NewMovieHandler will create new object
func NewMovieHandler(e *echo.Echo, m usecase.MovieUsecase) {
	handler := &movieHandler{
		movieUsecase: m,
	}

	e.GET("/movies", handler.Fetch)
	e.GET("/movies/:id", handler.GetByID)
}

func (h *movieHandler) Fetch(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	request := &domain.SearchRequest{
		Search: c.QueryParam("search"),
		Page:   page,
	}

	movies, err := h.movieUsecase.Fetch(request)
	if err != nil {
		return err
	}

	return c.JSON(200, movies)
}

func (h *movieHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	movie, err := h.movieUsecase.GetByID(id)
	if err != nil {
		return err
	}

	return c.JSON(200, movie)
}
