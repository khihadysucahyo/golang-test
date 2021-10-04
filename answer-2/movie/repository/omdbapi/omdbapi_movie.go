package omdbapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khihadysucahyo/golang-test/answer-2/domain"
)

type omdbapiMovieRepository struct {
	APIUrl string
}

// NewOmdbapiMovieRepository creates a new instance of omdbapiMovieRepository
func NewOmdbapiMovieRepository(APIKey string) domain.MovieRepository {
	return &omdbapiMovieRepository{APIUrl: "http://www.omdbapi.com/?apikey=" + APIKey}
}

func (r *omdbapiMovieRepository) Fetch(req *domain.SearchRequest) (res domain.SearchResponse, err error) {

	endpoint := r.APIUrl

	if req.Search != "" {
		endpoint += "&s=" + req.Search
	}

	if req.Page > 0 {
		endpoint += "&page=" + fmt.Sprintf("%d", req.Page)
	}

	httpRes, err := http.Get(endpoint)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	err = json.NewDecoder(httpRes.Body).Decode(&res)

	return
}

func (r *omdbapiMovieRepository) GetByID(id string) (res domain.DetailResult, err error) {
	endpoint := r.APIUrl + "&i=" + id
	httpRes, err := http.Get(endpoint)
	if err != nil {
		return domain.DetailResult{}, err
	}

	err = json.NewDecoder(httpRes.Body).Decode(&res)

	return
}
