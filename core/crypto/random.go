package crypto

type ICryptoRandomService interface {
	Int64() (int64, error)
	UInt64() (uint64, error)
}
