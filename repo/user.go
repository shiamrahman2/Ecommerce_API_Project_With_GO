package repo

import (
	"ecomerce/domain"
	"ecomerce/user"

	"github.com/jmoiron/sqlx"
)


type UserRepo interface {
   user.UserRepo   // Embedding
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(usr domain.User) (*domain.User, error) {

	query := `
	INSERT INTO users (
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	err := r.db.Get(&usr.ID, query,
		usr.FirstName,
		usr.LastName,
		usr.Email,
		usr.Password,
		usr.IsShopOwner,
	)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {

	var user domain.User

	query := `
	SELECT 
		id,
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1;
	`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
