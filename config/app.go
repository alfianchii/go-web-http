package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Address = fmt.Sprintf(":%s", GetENV("APP_PORT"))
	MongoClient *mongo.Client
	MongoDB *mongo.Database
	ExecTimeoutDuration = 10*time.Second
	TokenDuration = 1*time.Hour
	RedisClient *redis.Client
)

func InitENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetENV(key string) string {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error reading .env file")
	}

	return myEnv[key]
}

func dbConfig () string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", GetENV("DB_HOST"), GetENV("DB_PORT"), GetENV("DB_USERNAME"), GetENV("DB_DATABASE"))
}

func DBConnect() *sqlx.DB {
	db, err := sqlx.Open(GetENV("DB_CONNECTION"), dbConfig())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")
	return db
}

func mongoDBUri() string {
	return fmt.Sprintf("mongodb://%s:%s", GetENV("MONGODB_HOST"), GetENV("MONGODB_PORT"))
}

func InitMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoDBUri()).SetAuth(options.Credential{
		Username: GetENV("MONGODB_USERNAME"),
		Password: GetENV("MONGODB_PASSWORD"),
	})
	ctx, cancel := CtxTime()
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	
	fmt.Println("Successfully connected to the MongoDB!")
	MongoClient = client
	MongoDB = MongoClient.Database(GetENV("MONGODB_DATABASE"))
	return client
}

func CtxTime() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(CtxBg(), ExecTimeoutDuration)
	return ctx, cancel
}

func CtxBg() context.Context {
	return context.Background()
}

func InitRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: GetENV("REDIS_HOST") + ":" + GetENV("REDIS_PORT"),
		Username: GetENV("REDIS_USERNAME"),
		Password: GetENV("REDIS_PASSWORD"),
		DB: 0,
	})

	_, err := RedisClient.Ping(CtxBg()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	fmt.Println("Successfully connected to the Redis!")
	return RedisClient
}