package main

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	_logsearchRepo "github.com/khihadysucahyo/golang-test/answer-2/logsearch/repository/mysql"
	_movieGrpcDelivery "github.com/khihadysucahyo/golang-test/answer-2/movie/delivery/grpc"
	_movieHttpDelivery "github.com/khihadysucahyo/golang-test/answer-2/movie/delivery/http"
	_movieRepo "github.com/khihadysucahyo/golang-test/answer-2/movie/repository/omdbapi"
	_movieUcase "github.com/khihadysucahyo/golang-test/answer-2/movie/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _, err := sqlmock.New()
	if err != nil {
		fmt.Printf("failed to connect to database: %s", err)
	}
	defer db.Close()

	movieRepo := _movieRepo.NewOmdbapiMovieRepository("faf7e5bb")
	logRepo := _logsearchRepo.NewMysqlLogseacrhRepository(db)
	movieUcase := _movieUcase.NewMovieUsecase(movieRepo, logRepo)
	_movieHttpDelivery.NewMovieHandler(e, movieUcase)
	_movieGrpcDelivery.NewMovieHandler(movieUcase)

	log.Fatal(e.Start(":1323"))
}
