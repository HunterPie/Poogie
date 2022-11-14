package localization

import (
	"context"
	"github.com/Haato3o/poogie/core/cache"
	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/persistence/bucket"
	"strings"
)

const ChecksumsCacheKey = "LOCALIZATION_CHECKSUMS"

type Service struct {
	bucket      bucket.IBucket
	hashService crypto.IHashService
	cache       cache.ICache
}

func (s *Service) ListAvailableLocalizations(ctx context.Context) map[string]string {
	checksums, found := s.cache.Get(ChecksumsCacheKey)

	if found {
		return checksums.(map[string]string)
	}

	localizations := make(map[string]string, 0)

	for _, value := range s.bucket.FindAll(ctx) {
		filteredFileName := filterFileName(value)

		data, _ := s.bucket.FindBy(ctx, filteredFileName)

		checksum := s.hashService.Checksum(string(data))

		localizations[value] = checksum
	}

	s.cache.Set(ChecksumsCacheKey, localizations)

	return localizations
}

func NewService(
	bucket bucket.IBucket,
	hashService crypto.IHashService,
	cache cache.ICache,
) *Service {
	return &Service{
		bucket:      bucket,
		hashService: hashService,
		cache:       cache,
	}
}

func filterFileName(fileName string) string {
	noSuffix := strings.TrimSuffix(fileName, ".xml")
	return strings.TrimPrefix(noSuffix, "localization/")
}
