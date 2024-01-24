package domain

type Cart struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	ProductID int     `json:"product_id" gorm:"type:int"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	UserID    int     `json:"user_id" gorm:"type:int"`
	User      User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Quantity  int     `json:"quantity" gorm:"type:int"`
	Total     int     `json:"total" gorm:"type:int"`
}

func (c *Cart) TableName() string {
	return "carts"
}

type CartUsecase interface {
	GetAll(params map[string]any) ([]*Cart, error)
	Insert(cart *Cart) (*Cart, error)
	Update(id int, quantity int) (*Cart, error)
	Delete(id int) error
}

type CartRepository interface {
	FindAll(params map[string]any) ([]*Cart, error)
	FindByID(id int) (*Cart, error)
	Create(cart *Cart) (*Cart, error)
	Update(cart *Cart) (*Cart, error)
	Delete(id int) error
}
