package application

import domainRepository "github.com/Luis1f2/empleados_api/src/roles/domain/repository"
import "github.com/Luis1f2/empleados_api/src/roles/domain/entities"

type GetRoleByID struct {
	repo domainRepository.RolRepository
}

func NewGetRoleByID(repo domainRepository.RolRepository) *GetRoleByID {
	return &GetRoleByID{repo: repo}
}

func (g *GetRoleByID) Execute(id int) (*entities.Rol, error) {
	return g.repo.GetByID(id)
}