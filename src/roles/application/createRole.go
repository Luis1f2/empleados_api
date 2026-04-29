package application

import (
	"errors"
	"strings"

	domainRepository "github.com/Luis1f2/empleados_api/src/roles/domain/repository"
	"github.com/Luis1f2/empleados_api/src/roles/domain/entities"
)

type CreateRole struct {
	repo domainRepository.RolRepository
}

func NewCreateRole(repo domainRepository.RolRepository) *CreateRole {
	return &CreateRole{repo: repo}
}

func (c *CreateRole) Execute(nombre string) (*entities.Rol, error) {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return nil, errors.New("el nombre del rol es obligatorio")
	}

	_, err := c.repo.GetByName(nombre)
	if err == nil {
		return nil, errors.New("el rol ya existe")
	}

	return c.repo.Create(nombre)
}