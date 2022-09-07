package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrInvalidId        = errors.New("invalid user id")
	ErrFailedToFindUser = errors.New("failed to find user")
)

const ACCOUNTS_COLLECTION_NAME = "accounts"

type AccountBadgeSchema struct {
	Id        string    `bson:"id"`
	CreatedAt time.Time `bson:"created_at"`
}

type HuntStatisticsSummarySchema struct {
	Id        string    `bson:"id"`
	CreatedAt time.Time `bson:"created_at"`
}

type AccountSchema struct {
	Id                      primitive.ObjectID            `bson:"_id"`
	Username                string                        `bson:"username"`
	Password                string                        `bson:"password"`
	Email                   string                        `bson:"email"`
	Experience              int64                         `bson:"experience"`
	Rating                  int64                         `bson:"rating"`
	ClientId                string                        `bson:"client_id"`
	AvatarUrl               string                        `bson:"avatar_url"`
	Badges                  []AccountBadgeSchema          `bson:"badges"`
	HuntStatisticsSummaries []HuntStatisticsSummarySchema `bson:"hunt_statistics"`
	IsSupporter             bool                          `bson:"is_supporter"`
	CreatedAt               time.Time                     `bson:"created_at"`
	UpdatedAt               time.Time                     `bson:"updated_at"`
	LastSessionAt           time.Time                     `bson:"last_session_at"`
	IsArchived              bool                          `bson:"is_archived"`
	IsActive                bool                          `bson:"is_active"`
}

func toBadgeModels(badges []AccountBadgeSchema) []account.AccountBadgesModel {
	var models = make([]account.AccountBadgesModel, len(badges))

	for _, badge := range badges {
		models = append(models, account.AccountBadgesModel{
			Id:        badge.Id,
			CreatedAt: badge.CreatedAt,
		})
	}

	return models
}

func (schema AccountSchema) toAccountModel() account.AccountModel {
	return account.AccountModel{
		Id:                         schema.Id.Hex(),
		Username:                   schema.Username,
		Password:                   schema.Password,
		Email:                      schema.Email,
		ClientId:                   schema.ClientId,
		Experience:                 schema.Experience,
		Rating:                     schema.Rating,
		AvatarUri:                  schema.AvatarUrl,
		Badges:                     toBadgeModels(schema.Badges),
		HuntStatisticsSummaryModel: []account.HuntStatisticsSummaryModel{},
		IsSupporter:                schema.IsSupporter,
		CreatedAt:                  schema.CreatedAt,
		UpdatedAt:                  schema.UpdatedAt,
		LastSessionAt:              schema.LastSessionAt,
		IsArchived:                 schema.IsArchived,
	}
}

type AccountMongoRepository struct {
	*mongo.Collection
}

// AreCredentialsValid implements account.IAccountRepository
func (r *AccountMongoRepository) AreCredentialsValid(ctx context.Context, username string, password string) bool {
	query := bson.M{
		"username": username,
		"password": password,
	}

	var schema AccountSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil && err != mongo.ErrNoDocuments {
		return false
	}

	return err != mongo.ErrNoDocuments
}

// Create implements account.IAccountRepository
func (r *AccountMongoRepository) Create(ctx context.Context, model account.AccountModel) (account.AccountModel, error) {
	// TODO: Create index for Username and Email

	schema := AccountSchema{
		Id:            primitive.NewObjectID(),
		Username:      model.Username,
		Password:      model.Password,
		Email:         model.Email,
		ClientId:      model.ClientId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastSessionAt: time.Now(),
	}

	_, err := r.InsertOne(ctx, schema)

	if err != nil {
		log.Error("failed to create account", err)
		return model, account.ErrFailedToCreateAccount
	}

	return model, nil
}

// GetByUsername implements account.IAccountRepository
func (r *AccountMongoRepository) GetByUsername(ctx context.Context, username string) (account.AccountModel, error) {
	query := bson.M{
		"username": username,
	}

	var schema AccountSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil {
		return account.AccountModel{}, ErrFailedToFindUser
	}

	return schema.toAccountModel(), nil
}

// DeleteBy implements account.IAccountRepository
func (*AccountMongoRepository) DeleteBy(ctx context.Context, userId string) account.AccountModel {
	panic("unimplemented")
}

// GetById implements account.IAccountRepository
func (r *AccountMongoRepository) GetById(ctx context.Context, userId string) (account.AccountModel, error) {
	if !primitive.IsValidObjectID(userId) {
		return account.AccountModel{}, nil
	}

	id, _ := primitive.ObjectIDFromHex(userId)

	query := bson.M{
		"_id": id,
	}

	var schema AccountSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	if err != nil {
		return account.AccountModel{}, ErrFailedToFindUser
	}

	return schema.toAccountModel(), nil
}

// IsEmailTaken implements account.IAccountRepository
func (r *AccountMongoRepository) IsEmailTaken(ctx context.Context, email string) bool {
	query := bson.M{
		"email": email,
	}

	var schema AccountSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	return err != mongo.ErrNoDocuments
}

// IsUsernameTaken implements account.IAccountRepository
func (r *AccountMongoRepository) IsUsernameTaken(ctx context.Context, username string) bool {
	query := bson.M{
		"username": username,
	}

	var schema AccountSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	return err != mongo.ErrNoDocuments
}

// UpdateAvatar implements account.IAccountRepository
func (*AccountMongoRepository) UpdateAvatar(ctx context.Context, userId string, avatar string) account.AccountModel {
	panic("unimplemented")
}

// UpdatePassword implements account.IAccountRepository
func (*AccountMongoRepository) UpdatePassword(ctx context.Context, userId string, password string) account.AccountModel {
	panic("unimplemented")
}

func NewAccountRepository(db *mongo.Database) *AccountMongoRepository {
	return &AccountMongoRepository{db.Collection(ACCOUNTS_COLLECTION_NAME)}
}
