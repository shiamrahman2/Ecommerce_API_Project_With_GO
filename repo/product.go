package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)
type Product struct {
	ID          int     `json:"id" db:"id"`
	Tittle      string  `json:"tittle" db:"tittle"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"imageURL" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	List() ([]*Product, error)
	Get(productId int) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(productId int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db:db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {

	query := `
	INSERT INTO products(
		tittle,
		description,
		price,
		image_url
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)
	RETURNING id
	`

	row := r.db.QueryRow(
		query,
		p.Tittle,
		p.Description,
		p.Price,
		p.ImgURL,
	)

	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
func (r *productRepo) Get(productId int) (*Product, error) {
	var prd Product
	query:=`
     SELECT
	  id,
	  tittle,
	  description,
	  price,
	  image_url
	 FROM products
	 WHERE id=$1
	`
	 err:=r.db.Get(&prd,query,productId)
	 if err!=nil{
		if err==sql.ErrNoRows{
			return nil,nil
		}else{
			return nil,err
		}
	 }
	 return &prd,nil
}
func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product
	query:=`
     SELECT
	  id,
	  tittle,
	  description,
	  price,
	  image_url
	 FROM products
	`
	 err:=r.db.Select(&productList,query)
	 if err!=nil{
		return nil,err
	 }
	 return productList,nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	query:=`
	UPDATE products
	SET
	tittle=$1,
	description=$2,
	price=$3,
	image_url=$4
	WHERE id=$5
	`
	row:=r.db.QueryRow(
		query,
		product.Tittle,
		product.Description,
		product.Price,
		product.ImgURL,
	    product.ID,
	   )
	err:=row.Err()
	if err!=nil{
		return nil,err
	}
	return &product,nil
}
func (r *productRepo) Delete(productId int) error {
	query:=`
	DELETE FROM products WHERE id=$1
	`
	_ , err:=r.db.Exec(query,productId)
	if err!=nil{
		return err
	}
	return nil
}



