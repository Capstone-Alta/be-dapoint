package entities

import (
	"gorm.io/gorm"
)

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"required"`
}

type User struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
	Name         string `json:"name" form:"name"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Photo        string
	Role         string `gorm:"default:user"`
	TotalPoint   int
	Transactions []Transaction
	UserVouchers []UserVoucher
}

func ObjUser(dataName string, dataEmail, dataPassword string) (user *User) {
	return &User{
		Name:     dataName,
		Email:    dataEmail,
		Password: dataPassword,
	}
}

type UserRepository interface {
	FindById(id uint64) (user User, err error)
	FindAll() (users []User, err error)
	FindByQuery(key string, value interface{}) (user User, err error)
	Insert(data User) (id uint64, err error)
	Update(id int, data User) (user User, err error)
	PointUpdate(id int, data User) (ok bool, err error)
}

type UserService interface {
	GetById(id uint64) (user User, err error)
	GetAll() (users []User, err error)
	Create(data User) (id uint64, err error)
	Modify(id int, data User) (user User, err error)
	Login(data UserLogin) (user User, ok bool, err error)
	PointModify(id int, data User) (ok bool, err error)
}
