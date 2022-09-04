package account

import (
	"context"
	"errors"
	"time"

	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/persistence/account"
)

const (
	DefaultAvatarUri = "" // TODO: Add default profile picture
)

var (
	ErrAccountWithEmailAlreadyExists = errors.New("there's already an account associated to that email")
	ErrUsernameTaken                 = errors.New("username is taken")
	ErrWrongPassword                 = errors.New("invalid password")
	ErrAccountDoesNotExist           = errors.New("account does not exist")
)

type AccountService struct {
	repository    account.IAccountRepository
	cryptoService crypto.ICryptographyService
	hashService   crypto.IHashService
}

func (s *AccountService) CreateNewAccount(
	ctx context.Context,
	data AccountCreationRequest,
	clientId string,
) (account.AccountModel, error) {
	encryptedEmail := s.cryptoService.Encrypt(data.Email)

	if s.repository.IsEmailTaken(ctx, encryptedEmail) {
		return account.AccountModel{}, ErrAccountWithEmailAlreadyExists
	}

	if s.repository.IsUsernameTaken(ctx, data.Username) {
		return account.AccountModel{}, ErrUsernameTaken
	}

	hashedPassword := s.hashService.Hash(data.Password)

	model, err := s.repository.Create(ctx, account.AccountModel{
		Username:                   data.Username,
		Password:                   hashedPassword,
		Email:                      encryptedEmail,
		ClientId:                   clientId,
		AvatarUri:                  DefaultAvatarUri,
		Badges:                     make([]account.AccountBadgesModel, 0),
		HuntStatisticsSummaryModel: make([]account.HuntStatisticsSummaryModel, 0),
		IsSupporter:                false,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
		LastSessionAt:              time.Now(),
		IsArchived:                 false,
	})

	if err != nil {
		return model, err
	}

	return model, nil
}

func (s *AccountService) GetAccountById(ctx context.Context, userId string) (account.AccountModel, error) {
	user, err := s.repository.GetById(ctx, userId)

	if err != nil {
		return account.AccountModel{}, ErrAccountDoesNotExist
	}

	return user, nil
}

func (s *AccountService) UpdateAvatar(ctx context.Context, userId string, data AvatarUpdateRequest) account.AccountModel {
	return s.repository.UpdateAvatar(ctx, userId, data.AvatarUrl)
}
