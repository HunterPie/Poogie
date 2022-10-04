package bucket

type IBucket interface {
	FindMostRecent() (string, error)
	FindBy(name string) ([]byte, error)
	Upload(name string, data []byte) (bool, error)
}
