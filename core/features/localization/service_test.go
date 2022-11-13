package localization_test

import (
	"context"
	"github.com/Haato3o/poogie/core/features/localization"
	"github.com/Haato3o/poogie/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type ListLocalizationTestCase struct {
	Filename   string
	PrettyName string
	Content    string
	Hash       string
}

func TestService_ListAvailableLocalizations(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()

	ctx := context.TODO()
	bucket := mocks.NewMockIBucket(mocker)
	hashService := mocks.NewMockIHashService(mocker)
	cache := mocks.NewMockICache(mocker)

	service := localization.NewService(
		bucket,
		hashService,
		cache,
	)

	t.Run("should calculate file hashes from storage if it isn't in cache", func(t *testing.T) {
		files := []ListLocalizationTestCase{
			{
				Filename:   "localization/test1.xml",
				PrettyName: "test1",
				Content:    "test1_content",
				Hash:       "test1_hash",
			},
			{
				Filename:   "localization/test2.xml",
				PrettyName: "test2",
				Content:    "test2_content",
				Hash:       "test2_hash",
			},
			{
				Filename:   "localization/test3.xml",
				PrettyName: "test3",
				Content:    "test3_content",
				Hash:       "test3_hash",
			},
			{
				Filename:   "localization/test4.xml",
				PrettyName: "test4",
				Content:    "test4_content",
				Hash:       "test4_hash",
			},
		}
		filenames := make([]string, 0)

		for _, testCase := range files {
			filenames = append(filenames, testCase.Filename)
		}

		expected := map[string]string{
			"localization/test1.xml": "test1_hash",
			"localization/test2.xml": "test2_hash",
			"localization/test3.xml": "test3_hash",
			"localization/test4.xml": "test4_hash",
		}

		cache.EXPECT().
			Get(localization.ChecksumsCacheKey).
			Return(struct{}{}, false).
			Times(1)

		bucket.EXPECT().
			FindAll(ctx).
			Return(filenames).
			Times(1)

		for _, testCase := range files {
			bucket.EXPECT().
				FindBy(ctx, testCase.PrettyName).
				Return([]byte(testCase.Content), nil).
				Times(1)

			hashService.EXPECT().
				Hash(testCase.Content).
				Return(testCase.Hash).
				Times(1)
		}

		cache.EXPECT().
			Set(localization.ChecksumsCacheKey, expected).
			Times(1)

		actual := service.ListAvailableLocalizations(ctx)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %s, expected %s", actual, expected)
		}
	})

	t.Run("should return from cache if exists", func(t *testing.T) {
		expected := map[string]string{
			"localization/test1.xml": "test1_hash",
			"localization/test2.xml": "test2_hash",
			"localization/test3.xml": "test3_hash",
			"localization/test4.xml": "test4_hash",
		}

		cache.EXPECT().
			Get(localization.ChecksumsCacheKey).
			Return(expected, true).
			Times(1)

		bucket.EXPECT().
			FindAll(ctx).
			Times(0)

		bucket.EXPECT().
			FindBy(ctx, gomock.Any()).
			Times(0)

		hashService.EXPECT().
			Hash(gomock.Any()).
			Times(0)

		cache.EXPECT().
			Set(localization.ChecksumsCacheKey, gomock.Any()).
			Times(0)

		actual := service.ListAvailableLocalizations(ctx)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %s, expected %s", actual, expected)
		}
	})
}
