package repo

type Product struct {
	ID          int     `json:"id"`
	Tittle      string  `json:"tittle"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageURL"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	List() ([]*Product, error)
	Get(productId int) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(productId int) error
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	GenerateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}
	return nil, nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, pro := range r.productList {
		if pro.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}
func (r *productRepo) Delete(productId int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}

func GenerateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Tittle:      "Orange",
		Description: "Orange is Red.I Love Orange",
		Price:       100,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRe_ACiJ6hVHQxMGOLRo7xqzrF36Iy7AQBhzw&s",
	}
	prd2 := &Product{
		ID:          2,
		Tittle:      "Apple",
		Description: "Apple is Green.I Love Apple",
		Price:       40,
		ImgURL:      "https://thumbs.dreamstime.com/b/fresh-green-apple-26821143.jpg",
	}
	prd3 := &Product{
		ID:          3,
		Tittle:      "Banana",
		Description: "Banana is Green.I Love Eating Banana.It's provide instant energy",
		Price:       10,
		ImgURL:      "https://t4.ftcdn.net/jpg/14/38/40/45/360_F_1438404595_1xovTEsBNaKfKJAJ7bEjgHWfbglx2QF3.jpg",
	}
	prd4 := &Product{
		ID:          4,
		Tittle:      "Mango",
		Description: "Mango is my favourite Fruits.",
		Price:       30,
		ImgURL:      "https://c8.alamy.com/comp/2APAH3J/mango-fruits-tree-hanging-branch-chiang-mai-thailand-2APAH3J.jpg",
	}
	prd5 := &Product{
		ID:          5,
		Tittle:      "Jack-fruit",
		Description: "Jack-Fruit is boring.",
		Price:       70,
		ImgURL:      "https://png.pngtree.com/thumb_back/fh260/background/20241025/pngtree-ripe-jackfruit-on-a-tree-image_16379412.jpg",
	}
	prd6 := &Product{
		ID:          6,
		Tittle:      "Lichi",
		Description: "I Love Eating Lichi.",
		Price:       120,
		ImgURL:      "https://thumbs.dreamstime.com/b/litchi-fruits-17476031.jpg",
	}
	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)
	r.productList = append(r.productList, prd6)

}
