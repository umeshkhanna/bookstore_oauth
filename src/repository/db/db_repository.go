package db

import (
	"bookstore_oauth/src/domain/access_token"
	"bookstore_oauth/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Database Connection not implemented yet")
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
