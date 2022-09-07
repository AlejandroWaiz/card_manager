package firestore_adapter

import (
	"card_manager/src/model"
	"context"
	"fmt"
	"os"
)

func (db *DB) SaveCardIntoDatabase(thisCard model.Card) error {

	defer db.client.Close()

	ctx := context.Background()

	_, _, err := db.client.Collection(os.Getenv("Firestore-Collection")).Add(ctx, thisCard)

	if err != nil {
		err = fmt.Errorf("Err saving %v card: %v", thisCard.Name, err)
		return err
	}

	return nil
}
