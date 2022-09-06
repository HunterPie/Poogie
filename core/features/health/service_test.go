package health_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Haato3o/poogie/core/features/health"
	"github.com/Haato3o/poogie/mocks"
	"github.com/golang/mock/gomock"
)

func TestHealthService(t *testing.T) {
	mocker := gomock.NewController(t)
	defer mocker.Finish()

	database := mocks.NewMockIDatabase(mocker)
	ctx := context.TODO()

	database.EXPECT().
		IsHealthy(ctx).
		Return(true, nil).
		Times(1)

	service := health.NewService(database)

	t.Run("IsHealthy should return true when every service is healthy", func(t *testing.T) {
		actual, err := service.IsHealthy(ctx)

		if actual != true {
			t.Errorf("got %t, expected %t", actual, true)
		}

		if err != nil {
			t.Errorf("got %s, expected nil", err)
		}
	})

	errMockError := errors.New("mock error")

	database.EXPECT().
		IsHealthy(ctx).
		Return(false, errMockError).
		Times(1)

	t.Run("IsHealthy should return false and an error when any service is not healthy", func(t *testing.T) {
		actual, err := service.IsHealthy(ctx)

		if actual != false {
			t.Errorf("got %t, expected %t", actual, false)
		}

		if err != errMockError {
			t.Errorf("got %s, expected %s", err, errMockError)
		}
	})
}
