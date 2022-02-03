package db

import (
	"database/sql"
	"os"
	// "github.com/go-redis/redis/v8"
    _ "github.com/lib/pq"
)

// func InitRedis() *redis.Client {
// 	redisHost, redisPassword := os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD")
	
// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr: redisHost,
// 		Password: redisPassword,
// 		DB: 0,
// 	}) 

// 	return redisClient
// }


func ConnectDB() (*sql.DB, error) {
	userName, password, port, dbName := os.Getenv("POSTGRE_DB_USERNAME"), os.Getenv("POSTGRE_DB_PASSWORD"), os.Getenv("POSTGRE_DB_PORT"), os.Getenv("POSTGRE_DB_NAME")

	dbConfigs := "user="+userName+" dbname="+dbName+" password="+password+" host=localhost port="+port+" sslmode=disable"
	// dbURL := os.Getenv("DATABASE_URL")
	db, dbErr := sql.Open("postgres", dbConfigs)
	if dbErr != nil {
		return nil, dbErr
	}

	dbErr = db.Ping()
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
	
}