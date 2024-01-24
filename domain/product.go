package domain

type Product struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name" gorm:"type:varchar(255)"`
	Description string   `json:"description" gorm:"type:varchar(255)"`
	Price       int      `json:"price" gorm:"type:int"`
	Stock       int      `json:"stock" gorm:"type:int"`
	CategoryID  int      `json:"category_id" gorm:"type:int"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductUsecase interface {
	GetAll(params map[string]any) ([]*Product, error)
	GetByID(id int) (*Product, error)
	Create(product Product) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(id int) error
}

type ProductRepository interface {
	FindAll(params map[string]any) ([]*Product, error)
	FindByID(id int) (*Product, error)
	Create(product *Product) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id int) error
}
