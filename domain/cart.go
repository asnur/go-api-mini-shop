package domain

type Cart struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	ProductID int     `json:"product_id" gorm:"type:int"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    int     `json:"user_id" gorm:"type:int"`
	User      User    `gorm:"foreignKey:UserID"`
	Quantity  int     `json:"quantity" gorm:"type:int"`
	Total     int     `json:"total" gorm:"type:int"`
}

func (c *Cart) TableName() string {
	return "carts"
}

type CartUsecase interface {
	GetAll(user_id int) ([]*Cart, error)
	Insert(product_id int, user_id int, quantity int) (*Cart, error)
	Update(id int, quantity int) (*Cart, error)
	Delete(id int) error
}

type CartRepository interface {
	FindAll(user_id int) ([]*Cart, error)
	Create(cart *Cart) (*Cart, error)
	Update(cart *Cart) (*Cart, error)
	Delete(id int) error
}
