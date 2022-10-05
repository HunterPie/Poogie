package supporter

import (
	"context"

	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/email"
	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/supporter"
	"github.com/Haato3o/poogie/core/services"
)

const (
	SUPPORTER_EMAIL_TITLE = "HunterPie - Supporter Token"
)

type SupporterService struct {
	repository        supporter.ISupporterRepository
	emailService      email.IEmailService
	tokenService      services.TokenService
	accountRepository account.IAccountRepository
	cryptoService     crypto.ICryptographyService
}

func (s *SupporterService) CreateNewSupporter(ctx context.Context, email string) supporter.SupporterModel {

	var model supporter.SupporterModel

	if !s.repository.ExistsSupporter(ctx, email) {
		encryptedEmail := s.cryptoService.Encrypt(email)
		userAccount, _ := s.accountRepository.GetByEmail(ctx, encryptedEmail)

		userId := "TOKEN_NOT_ACTIVATED"

		if userAccount.Id != "" {
			s.accountRepository.UpdateSupporterStatus(ctx, userAccount.Id, true)
			userId = userAccount.Id
		}

		token := s.tokenService.Generate()
		model = s.repository.Insert(ctx, supporter.SupporterModel{
			UserId:   userId,
			Email:    email,
			Token:    token,
			IsActive: true,
		})

	} else {
		model = s.repository.RenewBy(ctx, email)

		if model.UserId != "TOKEN_NOT_ACTIVATED" {
			s.accountRepository.UpdateSupporterStatus(ctx, model.UserId, true)
		}
	}

	s.emailService.Send(
		SUPPORTER_EMAIL_TITLE,
		[]string{email},
		"./templates/supporter_token_email.html",
		struct {
			Token string
		}{
			Token: model.Token,
		})

	return model
}

func (s *SupporterService) RevokeExistingSupporter(ctx context.Context, email string) supporter.SupporterModel {
	model := s.repository.RevokeBy(ctx, email)

	if model.UserId != "TOKEN_NOT_ACTIVATED" {
		s.accountRepository.UpdateSupporterStatus(ctx, model.UserId, false)
	}

	return model
}

func (s *SupporterService) ExistsSupporterByToken(ctx context.Context, token string) bool {
	return s.repository.ExistsToken(ctx, token)
}
