package crypto

type ICryptographyService interface {
	Encrypt(content string) string
	Decrypt(content string) (string, error)
}
