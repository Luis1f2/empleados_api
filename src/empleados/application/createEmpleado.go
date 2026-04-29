package application

import (
	"errors"
	"strings"

	"github.com/Luis1f2/empleados_api/src/empleados/domain/entities"
	domainRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type CreateEmpleado struct {
	repo domainRepository.EmpleadoRepository
}

func NewCreateEmpleado(repo domainRepository.EmpleadoRepository) *CreateEmpleado {
	return &CreateEmpleado{repo: repo}
}

func (c *CreateEmpleado) Execute(
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

	if nombre == "" || apellidos == "" || mail == "" || user == "" || password == "" || rolID <= 0 {
		return nil, errors.New("todos los campos son obligatorios")
	}

	exists, err := c.repo.ExistsByUser(user)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("ya existe un empleado con ese username")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	empleado := &entities.Empleado{
		Nombre:    nombre,
		Apellidos: apellidos,
		Mail:      mail,
		User:      user,
		Password:  string(hashedPassword),
		RolID:     rolID,
	}

	return c.repo.Create(empleado)
}