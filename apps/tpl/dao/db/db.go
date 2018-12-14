package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql" //
	"github.com/jmoiron/sqlx"
	"github.com/xujintao/gourd/apps/tpl/dao"
	"github.com/xujintao/gourd/apps/tpl/dao/db/ddls"
)

type db struct {
	*sqlx.DB
}

// New creates a database connection for the given driver and datasource
// and returns a new Store.
func New(dsn string, maxConn int) dao.DB {
	return &db{
		DB: open(dsn, maxConn),
	}
}

func open(dsn string, maxConn int) *sqlx.DB {
	_db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	_db.SetMaxIdleConns(0)
	_db.SetMaxOpenConns(maxConn)

	// go func() {
	// 	if err := _db.Ping(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	log.Println("db connected.")
	// 	// setupDatabase(_db)
	// }()

	if err := _db.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("db connected.")

	if err := setupDatabase(_db); err != nil {
		log.Fatalln(err)
	}

	return _db
}

func setupDatabase(db *sqlx.DB) error {
	return ddls.Migrate(db.DB)
}
