package account

import (
	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
	"time"
)

type Account struct {
	Id           uuid.UUID   `json:"id" bson:"_id"`
	CreatedAt    time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt" bson:"updatedAt"`
	DeletedAt    *time.Time  `json:"deletedAt" bson:"deletedAt"`
	IsDeleted    bool        `json:"isDeleted" bson:"isDeleted"`
	Name         string      `json:"name" bson:"name"`
	Email        string      `json:"email" bson:"email"`
	CharacterIds []uuid.UUID `json:"characterIds" bson:"characterIds"`
}

func New(googleIdToken *idtoken.Payload) *Account {

	return &Account{
		Id:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		DeletedAt:    nil,
		IsDeleted:    false,
		Name:         googleIdToken.Claims["name"].(string),
		Email:        googleIdToken.Claims["email"].(string),
		CharacterIds: nil,
	}
}
