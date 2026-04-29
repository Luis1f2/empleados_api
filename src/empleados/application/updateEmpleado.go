package application

import (
	"errors"
	"strings"

	"github.com/Luis1f2/empleados_api/src/empleados/domain/entities"
	domainRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UpdateEmpleado struct {
	repo domainRepository.EmpleadoRepository
}

func NewUpdateEmpleado(repo domainRepository.EmpleadoRepository) *UpdateEmpleado {
	return &UpdateEmpleado{repo: repo}
}

func (u *UpdateEmpleado) Execute(
	id int,
	nombre string,
	apellidos string,
	mail string,
	user string,
	password string,
	rolID int,
) (*entities.Empleado, error) {
	nombre = strings.TrimSpace(nombre)
	apellidos = strings.TrimSpace(apellidos)
	mail = strings.TrimSpace(mail)
	user = strings.TrimSpace(user)
	password = strings.TrimSpace(password)

	if id <= 0 || nombre == "" || apellidos == "" || mail == "" || user == "" || rolID <= 0 {
		return nil, errors.New("datos inválidos")
	}

	actual, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	actual.Nombre = nombre
	actual.Apellidos = apellidos
	actual.Mail = mail
	actual.User = user
	actual.RolID = rolID

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		actual.Password = string(hashedPassword)
	}

	return u.repo.Update(id, actual)
}