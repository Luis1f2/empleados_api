package application

import domainRepository "github.com/Luis1f2/empleados_api/src/roles/domain/repository"
import "github.com/Luis1f2/empleados_api/src/roles/domain/entities"

type ListRoles struct {
	repo domainRepository.RolRepository
}

func NewListRoles(repo domainRepository.RolRepository) *ListRoles {
	return &ListRoles{repo: repo}
}

func (l *ListRoles) Execute() ([]entities.Rol, error) {
	return l.repo.List()
}