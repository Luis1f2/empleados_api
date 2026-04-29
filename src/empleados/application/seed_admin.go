package application

import (
	"fmt"

	"github.com/Luis1f2/empleados_api/src/empleados/domain/entities"
	empleadoRepository "github.com/Luis1f2/empleados_api/src/empleados/domain/repository"
	rolRepository "github.com/Luis1f2/empleados_api/src/roles/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type SeedAdminUseCase struct {
	EmpleadoRepo empleadoRepository.EmpleadoRepository
	RolRepo      rolRepository.RolRepository
}

func NewSeedAdminUseCase(
	empleadoRepo empleadoRepository.EmpleadoRepository,
	rolRepo rolRepository.RolRepository,
) *SeedAdminUseCase {
	return &SeedAdminUseCase{
		EmpleadoRepo: empleadoRepo,
		RolRepo:      rolRepo,
	}
}

func (s *SeedAdminUseCase) Execute() error {
	rol, err := s.RolRepo.GetByName("admin")
	if err != nil {
		rol, err = s.RolRepo.Create("admin")
		if err != nil {
			return fmt.Errorf("error al crear rol admin: %w", err)
		}
	}

	exists, err := s.EmpleadoRepo.ExistsByUser("admin")
	if err != nil {
		return fmt.Errorf("error al verificar empleado admin: %w", err)
	}

	if exists {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error al encriptar contraseña admin: %w", err)
	}

	admin := &entities.Empleado{
		Nombre:    "Administrador",
		Apellidos: "Sistema",
		Mail:      "admin@recolecta.com",
		User:      "admin",
		Password:  string(hashedPassword),
		RolID:     rol.ID,
	}

	_, err = s.EmpleadoRepo.Create(admin)
	if err != nil {
		return fmt.Errorf("error al crear empleado admin: %w", err)
	}

	return nil
}