package patches

import "context"

type Patch struct {
	Title       string
	Description string
	Link        string
	Banner      string
}

type IPatchRepository interface {
	FindAll(ctx context.Context) []Patch
}
