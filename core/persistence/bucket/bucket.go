package bucket

import (
	"context"
	"io"
)

type StreamedFile struct {
	Reader io.Reader
	Size   int64
	Type   string
}

type IBucket interface {
	FindMostRecent() (string, error)
	FindBy(name string) ([]byte, error)
	Upload(name string, data []byte) (bool, error)
	UploadFromStream(ctx context.Context, name string, file io.Reader) (bool, error)
	DownloadToStream(ctx context.Context, name string) (StreamedFile, error)
}
