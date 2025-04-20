package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:enal1233@tcp(localhost:3306)/belajar_golang_database_migration")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state
	// migrate -database "mysql://root:enal1233@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:enal1233@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down
	// migrate -database "mysql://root:enal1233@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version
	// migrate -database "mysql://root:enal1233@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20250420081052
}
