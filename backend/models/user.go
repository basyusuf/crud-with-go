package models

import (
	"errors"
	"main/app"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Status   bool   `json:"-" gorm:"default:true"`
}

type UserList []User

func (u *User) ToPublic() app.User {
	return app.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (list UserList) UserArrayToPublic() []app.User {
	s := make([]app.User, 0, 5)
	for _, value := range list {
		s = append(s, value.ToPublic())
	}
	return s
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(password)
	return
}
func (u *User) ValidateFor(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}
