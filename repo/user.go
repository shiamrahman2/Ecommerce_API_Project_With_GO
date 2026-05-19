package repo

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(usr User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(usr User) (*User, error) {

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

func (r *userRepo) Find(email, pass string) (*User, error) {

	var user User

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
