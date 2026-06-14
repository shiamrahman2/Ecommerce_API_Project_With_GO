package user

import(
	"ecomerce/domain"
)
type Service interface{
	Create(usr domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}