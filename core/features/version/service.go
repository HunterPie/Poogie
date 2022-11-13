package version

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/bucket"
	"github.com/Haato3o/poogie/core/persistence/patches"
)

type VersionService struct {
	bucket          bucket.IBucket
	alphaBucket     bucket.IBucket
	patchRepository patches.IPatchRepository
}

func (s *VersionService) GetLatestFileVersion(ctx context.Context, isSupporter bool) (string, error) {
	switch isSupporter {
	case true:
		return s.alphaBucket.FindMostRecent(ctx)
	default:
		return s.bucket.FindMostRecent(ctx)
	}
}

func (s *VersionService) GetFileByVersion(ctx context.Context, version string, isSupporter bool) ([]byte, error) {
	switch isSupporter {
	case true:
		return s.alphaBucket.FindBy(ctx, version)
	default:
		return s.bucket.FindBy(ctx, version)
	}
}

func (s *VersionService) GetPatchNotes(ctx context.Context) []patches.Patch {
	return s.patchRepository.FindAll(ctx)
}

func NewService(
	bucket bucket.IBucket,
	alphaBucket bucket.IBucket,
	patchRepository patches.IPatchRepository,
) *VersionService {
	return &VersionService{
		bucket:          bucket,
		alphaBucket:     alphaBucket,
		patchRepository: patchRepository,
	}
}
