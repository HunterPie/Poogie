package crypto

type IHashService interface {
	Hash(content string) string
	Checksum(content string) string
}
