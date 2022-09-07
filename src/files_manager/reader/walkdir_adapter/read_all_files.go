package walkdir_adapter

import (
	"card_manager/src/model"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var cardToCreate model.Card
var allErrors []string

func (r *ReaderImplementation) ReadAllFiles() error {

	godotenv.Load(".env")

	err := filepath.WalkDir(os.Getenv("Cards_Folder_Path"), walkDirFunc)

	if err != nil {
		xerr := fmt.Errorf("List of errors reading files: \n%v", strings.Join(allErrors, "\n"))
		return xerr
	}

	return nil

}

var walkDirFunc = func(path string, fileInfo fs.DirEntry, err error) error {

	if err != nil {
		obtainedErr := fmt.Sprintf("[ Path: %v | Step: WalkDir entry | Error: %v ]", path, err)
		allErrors = append(allErrors, obtainedErr)
		return err
	}

	if !fileInfo.IsDir() && strings.Contains(fileInfo.Name(), "JSON") {

		bytesToConvertToCardFormat, err := ioutil.ReadFile(path)

		if err != nil {
			obtainedErr := fmt.Sprintf("[ Path: %v | Step: Reading file | Error: %v ]", path, err)
			allErrors = append(allErrors, obtainedErr)
			return err
		}

		err = json.Unmarshal(bytesToConvertToCardFormat, &cardToCreate)

		if err != nil {
			obtainedErr := fmt.Sprintf("[ Path: %v | Step: Unmarshaling JSON | Error: %v ]", path, err)
			allErrors = append(allErrors, obtainedErr)
			return err
		}

	}

	return nil
}
