package repositories

import (
	"database/sql"
	"project-golang/models"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (repo *RoleRepository) GetAll() ([]models.Role, error) {
	query := `
			SELECT 
			id, 
			name FROM role`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	role := make([]models.Role, 0)
	for rows.Next() {
		var p models.Role
		err := rows.Scan(&p.ID, &p.Name)

		if err != nil {
			return nil, err
		}

		role = append(role, p)
	}

	return role, nil
}

func (repo *RoleRepository) Create(role *models.Role) error {
	query := `
			INSERT INTO role 
			(name) VALUES ($1) RETURNING id`
	err := repo.db.QueryRow(query, role.Name).Scan(&role.ID)
	return err
}
