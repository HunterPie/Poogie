package version

import (
	"context"

	"github.com/Haato3o/poogie/core/persistence/bucket"
	"github.com/Haato3o/poogie/core/persistence/patches"
	"github.com/Haato3o/poogie/core/persistence/supporter"
)

type VersionService struct {
	bucket              bucket.IBucket
	alphaBucket         bucket.IBucket
	supporterRepository supporter.ISupporterRepository
	patchRepository     patches.IPatchRepository
}

func (s *VersionService) GetLatestFileVersion(ctx context.Context, supporterToken string) (string, error) {
	isValidToken := s.supporterRepository.ExistsToken(ctx, supporterToken)

	switch isValidToken {
	case true:
		return s.alphaBucket.FindMostRecent(ctx)
	default:
		return s.bucket.FindMostRecent(ctx)
	}
}

func (s *VersionService) GetFileByVersion(ctx context.Context, version, supporterToken string) ([]byte, error) {
	isValidToken := s.supporterRepository.ExistsToken(ctx, supporterToken)

	switch isValidToken {
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
	supporterRepository supporter.ISupporterRepository,
	patchRepository patches.IPatchRepository,
) *VersionService {
	return &VersionService{
		bucket:              bucket,
		alphaBucket:         alphaBucket,
		supporterRepository: supporterRepository,
		patchRepository:     patchRepository,
	}
}
