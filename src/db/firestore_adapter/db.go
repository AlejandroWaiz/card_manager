package firestore_adapter

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type DB struct {
	client *firestore.Client
}

func GetFirestoreClient() (DB, error) {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"), option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))

	if err != nil {
		log.Printf("Err creating firestore client: %v", err)
		return DB{}, err
	}

	return DB{client: client}, nil

}

func (db *DB) CloseFirestoreClient() error {
	return db.client.Close()
}
