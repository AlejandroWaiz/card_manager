package db_port

import (
	"card_manager/src/db/firestore_adapter"
	"card_manager/src/model"
)

type DBInterface interface {
	SaveCardIntoDatabase(thisCard model.Card) error
	SaveAllCardsIntoDatabase(thisCards []model.Card) error
}

func GetFirestoreAdapter() (DBInterface, error) {

	firestoreImplementation, err := firestore_adapter.GetFirestoreClient()

	if err != nil {
		return nil, err
	}

	return &firestoreImplementation, nil

}
