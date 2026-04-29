package entities

import "time"

type Empleado struct {
	ID          int        `json:"id"`
	Nombre      string     `json:"nombre"`
	Apellidos   string     `json:"apellidos"`
	Mail        string     `json:"mail"`
	User        string     `json:"username"`
	Password    string     `json:"-"`
	RolID       int        `json:"rol_id"`
	Desactivado bool       `json:"desactivado"`
	Eliminado   *time.Time `json:"eliminado"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}