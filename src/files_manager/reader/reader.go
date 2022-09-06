package reader

type ReaderImplementation struct{}

func GetNewReader() ReaderInterface {

	return &ReaderImplementation{}

}

type ReaderInterface interface {
	ReadAllFiles() error
}
