package account

import (
	"context"
	"errors"
	"time"

	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/email"
	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/bucket"
	"github.com/Haato3o/poogie/core/persistence/supporter"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/log"
	"github.com/google/uuid"
)

const (
	DefaultAvatarUri       = "https://cdn.hunterpie.com/avatars/default.png"
	VerificationUri        = "https://api.hunterpie.com/v1/account/verify/"
	CDNAvatarsUri          = "https://cdn.hunterpie.com/avatars/"
	VerificationEmailTitle = "HunterPie - Account Verification"
)

var (
	ErrAccountWithEmailAlreadyExists = errors.New("there's already an account associated to that email")
	ErrUsernameTaken                 = errors.New("username is taken")
	ErrWrongPassword                 = errors.New("invalid password")
	ErrAccountDoesNotExist           = errors.New("account does not exist")
	ErrUnverifiedAccount             = errors.New("account is not verified")
	ErrAlreadyActivated              = errors.New("account already verified")
	ErrUnknownError                  = errors.New("something went wrong")
)

type AccountService struct {
	repository             account.IAccountRepository
	supporterRepository    supporter.ISupporterRepository
	cryptoService          crypto.ICryptographyService
	hashService            crypto.IHashService
	verificationRepository account.IAccountVerificationRepository
	emailService           email.IEmailService
	avatarStorage          bucket.IBucket
}

func (s *AccountService) VerifyAccount(ctx context.Context, token string) (bool, error) {
	userId, err := s.verificationRepository.Find(ctx, token)

	if err != nil {
		log.Error("error when verifying account", err)
		return false, ErrUnknownError
	}

	user, err := s.repository.GetById(ctx, userId)

	if err != nil {
		log.Error("error when verifying account", err)
		return false, ErrUnknownError
	}

	if user.IsActive {
		return false, ErrAlreadyActivated
	}

	s.repository.VerifyAccount(ctx, user.Id)

	return true, nil
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

	isSupporter := s.supporterRepository.ExistsSupporter(ctx, data.Email)

	model, err := s.repository.Create(ctx, account.AccountModel{
		Username:                   data.Username,
		Password:                   hashedPassword,
		Email:                      encryptedEmail,
		ClientId:                   clientId,
		AvatarUri:                  DefaultAvatarUri,
		Badges:                     make([]account.AccountBadgesModel, 0),
		HuntStatisticsSummaryModel: make([]account.HuntStatisticsSummaryModel, 0),
		IsSupporter:                isSupporter,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
		LastSessionAt:              time.Now(),
		IsArchived:                 false,
		IsActive:                   false,
	})

	if err != nil {
		return model, err
	}

	if isSupporter {
		s.supporterRepository.AssociateToUser(ctx, data.Email, model.Id)
	}

	model.Email, _ = s.cryptoService.Decrypt(model.Email)

	verificationToken := uuid.NewString()
	s.verificationRepository.Create(ctx, verificationToken, model.Id)

	s.emailService.Send(
		VerificationEmailTitle,
		[]string{data.Email},
		"./templates/activate_account_email.html",
		struct {
			Username string
			Link     string
		}{
			Username: data.Username,
			Link:     VerificationUri + verificationToken,
		},
	)

	return model, nil
}

func (s *AccountService) GetAccountById(ctx context.Context, userId string) (account.AccountModel, error) {
	user, err := s.repository.GetById(ctx, userId)

	if err != nil {
		return account.AccountModel{}, ErrAccountDoesNotExist
	}

	if !user.IsActive {
		return account.AccountModel{}, ErrUnverifiedAccount
	}

	return user, nil
}

func (s *AccountService) UpdateAvatar(ctx context.Context, userId string, avatar []byte) (account.AccountModel, error) {
	fileName := utils.NewRandomString() + ".png"

	ok, err := s.avatarStorage.Upload(fileName, avatar)

	if !ok || err != nil {
		return account.AccountModel{}, ErrUnknownError
	}

	account := s.repository.UpdateAvatar(ctx, userId, CDNAvatarsUri+fileName)
	account.Email, _ = s.cryptoService.Decrypt(account.Email)
	return account, nil
}
