package repositories

import (
	"database/sql"
	"errors"
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

func (repo *UserRepository) Create(user *models.User) error {
	query := `
			INSERT INTO "user"
			(username, password, role_id) VALUES ($1, $2, $3)
			RETURNING id`
	err := repo.db.QueryRow(query, user.Username, user.Password, user.RoleID).Scan(&user.ID)
	return err
}

func (repo *UserRepository) GetByID(id int) (*models.User, error) {
	query := `
			SELECT
			u.id,
			u.username,
			u.password,
			u.role_id,
			r.id,
			r.name
			FROM "user" u 
			JOIN role r ON r.id = u.role_id WHERE u.id=$1`

	var u models.User
	var r models.Role

	err := repo.db.QueryRow(query, id).Scan(
		&u.ID, &u.Username, &u.Password, &u.RoleID, &r.ID, &r.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	u.Role = &r

	return &u, nil
}

func (repo *UserRepository) Update(user *models.User) error {
	query := `
			UPDATE "user" SET
			username=$1,
			password=$2,
			role_id=$3
			WHERE id=$4`

	result, err := repo.db.Exec(query, user.Username, user.Password, user.RoleID, user.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("User not found!")
	}

	return nil
}

func (repo *UserRepository) Delete(id int) error {
	query := `
			DELETE FROM "user" 
			WHERE id=$1`
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found!")
	}

	return err
}

func (repo *UserRepository) GetByUsername(username string) (*models.User, error) {
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
		WHERE u.username = $1`

	var u models.User
	var r models.Role
	err := repo.db.QueryRow(query, username).
		Scan(&u.ID, &u.Username, &u.Password, &u.RoleID, &r.ID, &r.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	u.Role = &r
	return &u, nil
}
