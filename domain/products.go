package domain

type Product struct {
	ID          int     `json:"id" db:"id"`
	Tittle      string  `json:"tittle" db:"tittle"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"imageURL" db:"image_url"`
}