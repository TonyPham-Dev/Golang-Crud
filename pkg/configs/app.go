package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CloseDatabase(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ConnectDatabase(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, error := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, error

}

func PingDatabases(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connect success database")
	return nil
}

// config env
func Getenv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Cannot loading environment ")
	}
	return os.Getenv(key)
}
