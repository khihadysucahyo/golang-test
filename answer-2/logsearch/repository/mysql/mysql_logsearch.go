package mysql

import (
	"database/sql"
	"log"
	"time"

	"github.com/khihadysucahyo/golang-test/answer-2/domain"
)

type mysqlLogseacrhRepository struct {
	Conn *sql.DB
}

// NewMysqlLogseacrhRepository is mysqlLogseacrhRepository constructor
func NewMysqlLogseacrhRepository(Conn *sql.DB) domain.LogsearchRepository {
	return &mysqlLogseacrhRepository{Conn: Conn}
}

func (r *mysqlLogseacrhRepository) Store(searchWord string) (err error) {
	query := `INSERT logsearch SET searchword=?, time=?`
	stmt, err := r.Conn.Prepare(query)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = stmt.Exec(searchWord, time.Now().Format(time.RFC3339))
	return
}
