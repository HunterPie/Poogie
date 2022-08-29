package bucket

type FileBucket interface {
	FindMostRecent() (string, error)
	FindBy(name string) ([]byte, error)
}
