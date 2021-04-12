package database

import (
	"os"
	//"github.com/jackc/pgx/v4/pgxpool"
	//_ "github.com/lib/pq"
)

type Database struct{}

var host = os.Getenv("DB_HOST")
var password = os.Getenv("DB_NAME")
var user = os.Getenv("DB_USERNAME")
var dbname = os.Getenv("DB_PASSWORD")
var schema = os.Getenv("DB_SCHEMA")

//var db *pgxpool.Pool

func init() {
	//var err error
	//psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", user, password, host, dbname, schema)
	//db, err = pgxpool.Connect(context.Background(), psqlInfo)

	//if err != nil {
	//	panic(err)
	//}
}

//func (Database Database) GetDatabase() *pgxpool.Pool {
//	return db
//}
