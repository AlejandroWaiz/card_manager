package reader

import (
	"card_manager/src/model"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func (r *ReaderImplementation) ReadAllFiles() error {

	godotenv.Load(".env")

	pathOfMainCardsFolder := os.Getenv("Cards_Folder_Path")

	var cardToCreate model.Card

	err := filepath.WalkDir(pathOfMainCardsFolder, func(path string, fileInfo fs.DirEntry, err error) error {

		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !fileInfo.IsDir() && strings.Contains(fileInfo.Name(), "JSON") {

			bytesToConvertToCardFormat, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println(err)
			}

			err = json.Unmarshal(bytesToConvertToCardFormat, &cardToCreate)

			if err != nil {
				log.Println(err)
			}

			log.Printf("json: %v", cardToCreate)

		}

		if err != nil {
			log.Println(err)
		}

		return nil
	})

	if err != nil {
		log.Printf("[ Files_Manager | Reader | ERROR: %v ]", err)
		return err
	}

	return nil

}
