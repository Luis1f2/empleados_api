package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Luis1f2/empleados_api/src/config"

	// Roles
	rolesInfra "github.com/Luis1f2/empleados_api/src/roles/infrastructure"
	rolesRepo "github.com/Luis1f2/empleados_api/src/roles/infrastructure/repository"
	rolesRoutes "github.com/Luis1f2/empleados_api/src/roles/infrastructure/routes"

	// Empleados
	empleadosApp "github.com/Luis1f2/empleados_api/src/empleados/application"
	empleadosInfra "github.com/Luis1f2/empleados_api/src/empleados/infrastructure"
	empleadosRepo "github.com/Luis1f2/empleados_api/src/empleados/infrastructure/repository"
	empleadosRoutes "github.com/Luis1f2/empleados_api/src/empleados/infrastructure/routes"

	// Auth
	authController "github.com/Luis1f2/empleados_api/src/auth/infrastructure/controller"
	authRoutes "github.com/Luis1f2/empleados_api/src/auth/infrastructure/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	fmt.Println("Intentando conectar a la base de datos...")

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a PostgreSQL: ", err)
	}
	defer db.Close()

	fmt.Println("Conexión exitosa a PostgreSQL")

	// Seed automático del rol admin y empleado admin
	rolRepositoryForSeed := rolesRepo.NewRolRepositoryPostgres(db)
	empleadoRepositoryForSeed := empleadosRepo.NewEmpleadoRepositoryPostgres(db)

	seedAdmin := empleadosApp.NewSeedAdminUseCase(
		empleadoRepositoryForSeed,
		rolRepositoryForSeed,
	)

	err = seedAdmin.Execute()
	if err != nil {
		log.Fatal("Error al preparar admin inicial: ", err)
	}

	fmt.Println("Rol admin y empleado admin verificados correctamente")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Conexion con API exitosa",
		})
	})

	// Auth Login
	empleadoRepositoryForAuth := empleadosRepo.NewEmpleadoRepositoryPostgres(db)
	authControllerInstance := authController.NewAuthController(empleadoRepositoryForAuth)
	authRoutes.AuthRoutes(router, authControllerInstance)

	// Roles
	rolesDeps := rolesInfra.InitRolesDependencies(db)
	rolesRoutes.RolRoutes(router, rolesDeps.RolController)

	// Empleados
	empleadosDeps := empleadosInfra.InitEmpleadosDependencies(db)
	empleadosRoutes.EmpleadoRoutes(router, empleadosDeps.EmpleadoController)

	fmt.Println("Servidor API corriendo en http://localhost:8080")

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar la API: ", err)
	}
}