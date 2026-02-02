package repositories

import (
	"database/sql"
	"project-golang/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetAll() ([]models.User, error) {
	query := `
		SELECT 
			u.id,
			u.username,
			u.password,
			u.role_id,
			r.id,
			r.name
		FROM "user" u
		JOIN role r ON r.id = u.role_id
	`
	
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		var u models.User
		var r models.Role

		err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Password,
			&u.RoleID,
			&r.ID,
			&r.Name,
		)
		if err != nil {
			return nil, err
		}

		u.Role = &r
		users = append(users, u)
	}

	return users, nil
}
