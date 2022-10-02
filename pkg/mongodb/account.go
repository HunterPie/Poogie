package mongodb

import (
	"context"
	"errors"
	"strings"
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
	UsernameUnique          string                        `bson:"username_unique"`
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
		IsActive:                   schema.IsActive,
	}
}

func toAccountSchema(model account.AccountModel) AccountSchema {
	return AccountSchema{
		Id:                      primitive.NewObjectID(),
		UsernameUnique:          strings.ToLower(model.Username),
		Username:                model.Username,
		Password:                model.Password,
		Email:                   model.Email,
		Experience:              model.Experience,
		Rating:                  model.Rating,
		ClientId:                model.ClientId,
		AvatarUrl:               model.AvatarUri,
		Badges:                  []AccountBadgeSchema{},
		HuntStatisticsSummaries: []HuntStatisticsSummarySchema{},
		IsSupporter:             model.IsSupporter,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
		LastSessionAt:           time.Now(),
		IsArchived:              model.IsArchived,
		IsActive:                model.IsActive,
	}
}

type AccountMongoRepository struct {
	*mongo.Collection
}

// VerifyAccount implements account.IAccountRepository
func (r *AccountMongoRepository) VerifyAccount(ctx context.Context, userId string) {
	id, _ := primitive.ObjectIDFromHex(userId)

	query := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"is_active": true,
		},
	}

	var schema AccountSchema
	_ = r.FindOneAndUpdate(ctx, query, update).Decode(&schema)

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
	schema := toAccountSchema(model)

	_, err := r.InsertOne(ctx, schema)

	if err != nil {
		log.Error("failed to create account", err)
		return model, account.ErrFailedToCreateAccount
	}

	query := bson.M{
		"email": model.Email,
	}

	_ = r.FindOne(ctx, query).Decode(&schema)

	return schema.toAccountModel(), nil
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
		"username_unique": strings.ToLower(username),
	}

	var schema AccountSchema
	err := r.FindOne(ctx, query).Decode(&schema)

	return err != mongo.ErrNoDocuments
}

// UpdateAvatar implements account.IAccountRepository
func (r *AccountMongoRepository) UpdateAvatar(ctx context.Context, userId string, avatar string) account.AccountModel {
	id, _ := primitive.ObjectIDFromHex(userId)

	query := bson.M{
		"_id": id,
	}

	update := bson.M{
		"avatar_url": avatar,
	}

	var schema AccountSchema
	_ = r.FindOneAndUpdate(ctx, query, update).Decode(&schema)

	return schema.toAccountModel()
}

// UpdatePassword implements account.IAccountRepository
func (*AccountMongoRepository) UpdatePassword(ctx context.Context, userId string, password string) account.AccountModel {
	panic("unimplemented")
}

func NewAccountRepository(db *mongo.Database) *AccountMongoRepository {
	return &AccountMongoRepository{db.Collection(ACCOUNTS_COLLECTION_NAME)}
}
