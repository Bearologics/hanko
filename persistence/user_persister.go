package persistence

import (
	"database/sql"
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/teamhanko/hanko/models"
)

type UserPersister struct {
	db *pop.Connection
}

func NewUserPersister(db *pop.Connection) *UserPersister {
	return &UserPersister{db: db}
}

func (p *UserPersister) Get(id uuid.UUID) (*models.User, error) {
	user := models.User{}
	err := p.db.Find(&user, id)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (p *UserPersister) Create(user models.User) error {
	vErr, err := p.db.ValidateAndCreate(&user)
	if err != nil {
		return fmt.Errorf("failed to store user: %w", err)
	}

	if vErr != nil && vErr.HasAny() {
		return fmt.Errorf("user object validation failed: %w", vErr)
	}

	return nil
}

func (p *UserPersister) Update(user models.User) error {
	vErr, err := p.db.ValidateAndUpdate(&user)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if vErr != nil && vErr.HasAny() {
		return fmt.Errorf("user object validation failed: %w", vErr)
	}

	return nil
}

func (p *UserPersister) Delete(user models.User) error {
	err := p.db.Destroy(&user)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
