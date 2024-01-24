package domain

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}

func (u *User) TableName() string {
	return "users"
}

type UserUsecase interface {
	Login(user User) (string, error)
	Register(user User) (*User, error)
}

type UserRepository interface {
	FindUserByEmail(email string) (*User, error)
	FindByID(id int) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id int) error
}
