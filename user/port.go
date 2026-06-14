package user

import (
	"ecomerce/domain"
	userHandler "ecomerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service // Embedding
}

type UserRepo interface {
	Create(usr domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
