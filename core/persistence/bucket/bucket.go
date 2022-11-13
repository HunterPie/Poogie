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
	FindMostRecent(ctx context.Context) (string, error)
	FindBy(ctx context.Context, name string) ([]byte, error)
	Upload(ctx context.Context, name string, data []byte) (bool, error)
	UploadFromStream(ctx context.Context, name string, file io.Reader) (bool, error)
	DownloadToStream(ctx context.Context, name string) (StreamedFile, error)
	Delete(ctx context.Context, name string)
	FindAll(ctx context.Context) []string
}
