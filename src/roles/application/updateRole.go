package application

import (
	"errors"
	"strings"

	domainRepository "github.com/Luis1f2/empleados_api/src/roles/domain/repository"
	"github.com/Luis1f2/empleados_api/src/roles/domain/entities"
)

type UpdateRole struct {
	repo domainRepository.RolRepository
}

func NewUpdateRole(repo domainRepository.RolRepository) *UpdateRole {
	return &UpdateRole{repo: repo}
}

func (u *UpdateRole) Execute(id int, nombre string) (*entities.Rol, error) {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return nil, errors.New("el nombre del rol es obligatorio")
	}

	return u.repo.Update(id, nombre)
}