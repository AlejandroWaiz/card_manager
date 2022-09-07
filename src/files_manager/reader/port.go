package reader_port

import "card_manager/src/files_manager/reader/walkdir_adapter"

type ReaderInterface interface {
	ReadAllFiles() error
}

func GetWalkDirReaderAdapter() ReaderInterface {

	return &walkdir_adapter.ReaderImplementation{}

}
