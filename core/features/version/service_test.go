package version_test

import (
	"bytes"
	"context"
	"github.com/Haato3o/poogie/core/features/version"
	"github.com/Haato3o/poogie/core/persistence/patches"
	"github.com/Haato3o/poogie/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestVersionService_GetLatestFileVersion(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()

	ctx := context.TODO()
	bucket := mocks.NewMockIBucket(mocker)
	alphaBucket := mocks.NewMockIBucket(mocker)
	supporterRepository := mocks.NewMockISupporterRepository(mocker)
	patchRepository := mocks.NewMockIPatchRepository(mocker)

	service := version.NewService(
		bucket,
		alphaBucket,
		supporterRepository,
		patchRepository,
	)

	t.Run("should get latest version from normal bucket if user is not supporter", func(t *testing.T) {
		supporterToken := gofakeit.Word()
		mostRecentFile := gofakeit.Word()

		supporterRepository.EXPECT().
			ExistsToken(ctx, supporterToken).
			Return(false).
			Times(1)

		bucket.EXPECT().
			FindMostRecent(ctx).
			Return(mostRecentFile, nil).
			Times(1)

		alphaBucket.EXPECT().
			FindMostRecent(ctx).
			Times(0)

		actual, err := service.GetLatestFileVersion(ctx, supporterToken)

		if actual != mostRecentFile {
			t.Errorf("got %s, expected %s", actual, mostRecentFile)
		}

		if err != nil {
			t.Errorf("got %s, expected nil", err)
		}
	})

	t.Run("should get latest version from alpha bucket if user is supporter", func(t *testing.T) {
		supporterToken := gofakeit.Word()
		mostRecentFile := gofakeit.Word()

		supporterRepository.EXPECT().
			ExistsToken(ctx, supporterToken).
			Return(true).
			Times(1)

		alphaBucket.EXPECT().
			FindMostRecent(ctx).
			Return(mostRecentFile, nil).
			Times(1)

		bucket.EXPECT().
			FindMostRecent(ctx).
			Times(0)

		actual, err := service.GetLatestFileVersion(ctx, supporterToken)

		if actual != mostRecentFile {
			t.Errorf("got %s, expected %s", actual, mostRecentFile)
		}

		if err != nil {
			t.Errorf("got %s, expected nil", err)
		}
	})
}

func TestVersionService_GetFileByVersion(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()

	ctx := context.TODO()
	bucket := mocks.NewMockIBucket(mocker)
	alphaBucket := mocks.NewMockIBucket(mocker)
	supporterRepository := mocks.NewMockISupporterRepository(mocker)
	patchRepository := mocks.NewMockIPatchRepository(mocker)

	service := version.NewService(
		bucket,
		alphaBucket,
		supporterRepository,
		patchRepository,
	)

	t.Run("should get file from bucket if user is not supporter", func(t *testing.T) {
		supporterToken := gofakeit.Word()
		fileBytes := []byte(gofakeit.Phrase())
		versionString := gofakeit.Word()

		supporterRepository.EXPECT().
			ExistsToken(ctx, supporterToken).
			Return(false).
			Times(1)

		bucket.EXPECT().
			FindBy(ctx, versionString).
			Return(fileBytes, nil).
			Times(1)

		alphaBucket.EXPECT().
			FindBy(ctx, versionString).
			Times(0)

		actual, err := service.GetFileByVersion(ctx, versionString, supporterToken)

		if bytes.Compare(actual, fileBytes) != 0 {
			t.Errorf("got %s, expected %s", actual, fileBytes)
		}

		if err != nil {
			t.Errorf("got %s, expected nil", err)
		}
	})

	t.Run("should get file from bucket if user is not supporter", func(t *testing.T) {
		supporterToken := gofakeit.Word()
		fileBytes := []byte(gofakeit.Phrase())
		versionString := gofakeit.Word()

		supporterRepository.EXPECT().
			ExistsToken(ctx, supporterToken).
			Return(true).
			Times(1)

		alphaBucket.EXPECT().
			FindBy(ctx, versionString).
			Return(fileBytes, nil).
			Times(1)

		bucket.EXPECT().
			FindBy(ctx, versionString).
			Times(0)

		actual, err := service.GetFileByVersion(ctx, versionString, supporterToken)

		if bytes.Compare(actual, fileBytes) != 0 {
			t.Errorf("got %s, expected %s", actual, fileBytes)
		}

		if err != nil {
			t.Errorf("got %s, expected nil", err)
		}
	})
}

func TestVersionService_GetPatchNotes(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()

	ctx := context.TODO()
	bucket := mocks.NewMockIBucket(mocker)
	alphaBucket := mocks.NewMockIBucket(mocker)
	supporterRepository := mocks.NewMockISupporterRepository(mocker)
	patchRepository := mocks.NewMockIPatchRepository(mocker)

	service := version.NewService(
		bucket,
		alphaBucket,
		supporterRepository,
		patchRepository,
	)

	t.Run("should return patch array", func(t *testing.T) {
		patchArray := make([]patches.Patch, 5)
		mocks.RandomSlice(patchArray)

		patchRepository.EXPECT().
			FindAll(ctx).
			Return(patchArray).
			Times(1)

		actual := service.GetPatchNotes(ctx)

		if !reflect.DeepEqual(actual, patchArray) {
			t.Errorf("got %s, expected %s", actual, patchArray)
		}
	})
}
