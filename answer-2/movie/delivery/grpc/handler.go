package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/khihadysucahyo/golang-test/answer-2/domain"
	"github.com/khihadysucahyo/golang-test/answer-2/movie/delivery/grpc/pb"
	"github.com/khihadysucahyo/golang-test/answer-2/movie/usecase"
	"google.golang.org/grpc"
)

type movieService struct {
	movieUsecase usecase.MovieUsecase
}

// NewMovieHandler will create new object
func NewMovieHandler(m usecase.MovieUsecase) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterMovieServiceServer(s, &movieService{})

	go func() {
		fmt.Println("⇨ GRPC server started on [::]:50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve %v", err)
		}
	}()
}

func (m *movieService) Fetch(ctx context.Context, in *pb.SearchRequest) (res *pb.SearchResponse, err error) {
	page, _ := strconv.Atoi(in.Page)

	movies, err := m.movieUsecase.Fetch(&domain.SearchRequest{
		Search: in.Search,
		Page:   page,
	})

	if err != nil {
		return
	}

	lists := make([]*pb.Movie, 0)

	for _, movie := range movies.Search {
		lists = append(lists, &pb.Movie{
			Title: movie.Title,
			Year:  movie.Year,
		})
	}

	return &pb.SearchResponse{Movies: lists}, nil
}

func (m *movieService) GetByID(ctx context.Context, in *pb.DetailRequest) (res *pb.Movie, err error) {
	movie, err := m.movieUsecase.GetByID(in.Id)
	if err != nil {
		return
	}
	return &pb.Movie{
		Title:  movie.Title,
		Year:   movie.Year,
		ImdbID: movie.ImdbID,
		Type:   movie.Type,
		Poster: movie.Poster,
	}, nil
}
