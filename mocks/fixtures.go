package mocks

import "github.com/brianvoe/gofakeit/v6"

func RandomSlice[T any](slice []T) {
	for i := 0; i < cap(slice); i++ {
		var data T
		_ = gofakeit.Struct(&data)

		slice[i] = data
	}
}

func RandomString() string {
	return gofakeit.Word()
}
