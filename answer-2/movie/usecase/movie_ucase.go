package usecase

import (
	"github.com/khihadysucahyo/golang-test/answer-2/domain"
)

// MovieUsecase interface
type MovieUsecase interface {
	Fetch(req *domain.SearchRequest) (domain.SearchResponse, error)
	GetByID(id string) (domain.DetailResult, error)
}

type movieUsecase struct {
	movieRepo domain.MovieRepository
	logRepo   domain.LogsearchRepository
}

// NewMovieUsecase will create new an MovieUsecase object representation of movieUsecase
func NewMovieUsecase(movieRepo domain.MovieRepository, logRepo domain.LogsearchRepository) MovieUsecase {
	return &movieUsecase{
		movieRepo: movieRepo,
		logRepo:   logRepo,
	}
}

func (m *movieUsecase) Fetch(req *domain.SearchRequest) (res domain.SearchResponse, err error) {
	res, err = m.movieRepo.Fetch(req)
	m.logRepo.Store(req.Search)
	return
}

func (m *movieUsecase) GetByID(id string) (domain.DetailResult, error) {
	return m.movieRepo.GetByID(id)
}
