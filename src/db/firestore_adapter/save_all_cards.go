package firestore_adapter

import (
	"card_manager/src/model"
	"context"
	"fmt"
	"os"
	"strings"
)

func (db *DB) SaveAllCardsIntoDatabase(thisCards []model.Card) error {

	var unsavedCards []string

	ctx := context.Background()

	for _, card := range thisCards {

		_, _, err := db.client.Collection(os.Getenv("Firestore-Collection")).Add(ctx, card)

		if err != nil {

			unsavedCards = append(unsavedCards, card.Name)

		}

	}

	if len(unsavedCards) > 0 {

		err := fmt.Errorf("Error saving this cards:\n%v", strings.Join(unsavedCards, "\n"))

		return err

	}

	return nil

}
