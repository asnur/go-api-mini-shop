package domain

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	Slug string `json:"slug" gorm:"type:varchar(255)"`
}

func (c *Category) TableName() string {
	return "categories"
}

type CategoryUsecase interface {
	GetAll() ([]*Category, error)
	Detail(id int) (*Category, error)
	Insert(name string) (*Category, error)
	Update(id int, name string) (*Category, error)
	Delete(id int) error
}

type CategoryRepository interface {
	FindAll() ([]*Category, error)
	FindByID(id int) (*Category, error)
	Create(category *Category) (*Category, error)
	Update(category *Category) (*Category, error)
	Delete(id int) error
}
