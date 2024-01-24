package domain

type Order struct {
	ID           int           `json:"id" gorm:"primaryKey"`
	UserID       int           `json:"user_id" gorm:"type:int"`
	User         User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Total        int           `json:"total" gorm:"type:int"`
	Status       int           `json:"status" gorm:"type:int"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (o *Order) TableName() string {
	return "orders"
}

type OrderUsecase interface {
	GetAll(user_id int) ([]*Order, error)
	GetByID(id int) (*Order, error)
	Insert(user_id int, total int, status int) (*Order, error)
	Delete(id int) error
}

type OrderRepository interface {
	FindAll(user_id int) ([]*Order, error)
	FindByID(id int) (*Order, error)
	Create(order *Order) (*Order, error)
	Delete(id int) error
}

type OrderDetail struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	OrderID   int     `json:"order_id" gorm:"type:int;"`
	Order     Order   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID int     `json:"product_id" gorm:"type:int"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Quantity  int     `json:"quantity" gorm:"type:int"`
	Total     int     `json:"total" gorm:"type:int"`
}

func (o *OrderDetail) TableName() string {
	return "order_details"
}
