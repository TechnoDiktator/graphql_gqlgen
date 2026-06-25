package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConnection()(*pgxpool.Pool, error) {
	
	//example connection string: "postgres://postgres:1997@localhost:5432/graphql_demo"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_PASSWORD"), 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_PORT"), 
	os.Getenv("DB_NAME"))

	pool, err := pgxpool.New(context.Background(), connectionString)


	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return pool, nil
}




