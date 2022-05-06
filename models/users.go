package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/security"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) Prepare(stage string) error {

	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}
	return nil
}

func (user *User) format(stage string) error {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "create" {
		password, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(password)
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.FirstName == "" {
		return errors.New("O primeiro nome é obrigatório  e não pode estar em branco")
	}

	if user.LastName == "" {
		return errors.New("O sobrenome é obrigatório  e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatório  e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O email inserido é inválido!")
	}

	if stage == "create" && user.Password == "" {
		return errors.New("A senha é obrigatório  e não pode estar em branco")
	}

	return nil

}
