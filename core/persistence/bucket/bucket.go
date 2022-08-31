package bucket

type IBucket interface {
	FindMostRecent() (string, error)
	FindBy(name string) ([]byte, error)
}
