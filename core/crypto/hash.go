package crypto

type IHashService interface {
	Hash(content string) string
}
