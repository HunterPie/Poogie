package supporter

import (
	"context"

	"github.com/Haato3o/poogie/core/email"
	"github.com/Haato3o/poogie/core/persistence/supporter"
	"github.com/Haato3o/poogie/core/services"
)

const (
	SUPPORTER_EMAIL_TITLE = "HunterPie - Supporter Token"
)

type SupporterService struct {
	repository   supporter.ISupporterRepository
	emailService email.IEmailService
	tokenService services.TokenService
}

func (s *SupporterService) CreateNewSupporter(ctx context.Context, email string) supporter.SupporterModel {
	token := s.tokenService.Generate()

	model := s.repository.Insert(ctx, supporter.SupporterModel{
		Email:    email,
		Token:    token,
		IsActive: true,
	})

	s.emailService.Send(
		SUPPORTER_EMAIL_TITLE,
		[]string{email},
		"./templates/supporter_token_email.html",
		struct {
			Token string
		}{
			Token: token,
		})

	return model
}

func (s *SupporterService) RevokeExistingSupporter(ctx context.Context, email string) supporter.SupporterModel {
	return s.repository.RevokeBy(ctx, email)
}

func (s *SupporterService) ExistsSupporterByToken(ctx context.Context, token string) bool {
	return s.repository.ExistsToken(ctx, token)
}
