package domain

// LogsearchRepository is an interface that represents a log search repository
type LogsearchRepository interface {
	Store(searchWord string) error
}

// LogsearchUsecase is an interface that represents a log search service
type LogsearchUsecase interface {
	Store(searchWord string) error
}
