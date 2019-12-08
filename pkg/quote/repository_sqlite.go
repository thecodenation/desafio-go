package quote

import (
	"database/sql"

	"github.com/eucleciojosias/codenation-challenge/pkg/entity"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepository struct {
	db   *sql.DB
}

func NewSqliteRepository() *SqliteRepository {
	db, _ := sql.Open("sqlite3", "./database.sqlite")
	db.Exec("CREATE TABLE IF NOT EXISTS `scripts` (`index` INTEGER, `episode` INTEGER, `episode_name` TEXT," +
		"`segment` TEXT, `type` TEXT, `actor` TEXT, `character` TEXT, `detail` TEXT, `record_date` TIMESTAMP," +
		"`series` TEXT, `transmission_date` TIMESTAMP)")
	db.Exec("CREATE INDEX `ix_scripts_index` ON scripts (`index`)")

	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) FindByActor(actor string) ([]*entity.Quote, error) {
	rows, err := r.db.Query("SELECT detail, actor FROM scripts WHERE actor LIKE ?", actor)
	if err != nil {
		return nil, err
	}

	return scanQuotes(rows, err)
}

func (r *SqliteRepository) FindAll() ([]*entity.Quote, error) {
	rows, err := r.db.Query("SELECT detail, actor FROM scripts")
	if err != nil {
		return nil, err
	}

	return scanQuotes(rows, err)
}

func scanQuotes(rows *sql.Rows, err error) ([]*entity.Quote, error) {
	var quotes []*entity.Quote

	for rows.Next() {
		var detail string
		var actor string
		err := rows.Scan(&detail, &actor)
		if err != nil {
			return nil, err
		}
		var quote entity.Quote
		quote.Detail = detail
		quote.Actor = actor
		quotes = append(quotes, &quote)
	}

	return quotes, err
}
