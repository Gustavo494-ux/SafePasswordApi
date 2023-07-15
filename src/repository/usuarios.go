package repository

import (
	"errors"
	"safePasswordApi/src/models"

	"github.com/jmoiron/sqlx"
)

type Users struct {
	db *sqlx.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *sqlx.DB) *Users {
	return &Users{db}
}

// CreateUser adds a new user to the database.
func (repository Users) CreateUser(user models.Usuario) (uint64, error) {
	statement, err := repository.db.Exec(
		`INSERT INTO Users (nome, email, password) VALUES (?, ?, ?)`,
		user.Nome,
		user.Email,
		user.Senha,
	)
	rowsAffected, err := statement.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("no rows affected, check the provided data")
	}
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	userID, err := statement.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(userID), nil
}

// FindByID finds a user in the database by ID.
func (repository Users) FindByID(userID uint64) (models.Usuario, error) {
	user := models.Usuario{}
	err := repository.db.Get(&user,
		`SELECT id, nome, email, created_at FROM Users WHERE id = ?`,
		userID,
	)

	if user.ID == 0 {
		return models.Usuario{}, errors.New("no user found, check the provided data")
	}

	if err != nil {
		return models.Usuario{}, err
	}
	return user, nil
}

// FindAllUsers retrieves all users saved in the database.
func (repository Users) FindAllUsers() ([]models.Usuario, error) {
	var users []models.Usuario
	err := repository.db.Select(&users, "SELECT id, nome, email, password FROM Users")
	if len(users) == 0 {
		return []models.Usuario{}, errors.New("no users found, check the provided data")
	}

	if err != nil {
		return []models.Usuario{}, err
	}
	return users, nil
}

// UpdateUser updates user information in the database.
func (repository Users) UpdateUser(userID uint64, user models.Usuario) error {
	statement, err := repository.db.Exec(
		`UPDATE Users SET nome=?, email=?, password=? WHERE id=?`,
		user.Nome,
		user.Email,
		user.Senha,
		userID,
	)
	rowsAffected, err := statement.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no records affected, check the provided data")
	}

	if err != nil {
		return err
	}
	return nil
}

// FindByEmail finds a user by email and returns its ID and hashed password.
func (repository Users) FindByEmail(email string) (models.Usuario, error) {
	user := models.Usuario{}
	err := repository.db.Get(&user, "SELECT id, password FROM Users WHERE email = ?", email)
	if user.ID == 0 {
		return models.Usuario{}, errors.New("no user found, check the provided data")
	}
	if err != nil {
		return models.Usuario{}, err
	}
	return user, nil
}

// DeleteUser deletes a user from the database.
func (repository Users) DeleteUser(userID uint64) error {
	statement, err := repository.db.Exec(
		`DELETE FROM Users WHERE id = ?`,
		userID,
	)
	rowsAffected, err := statement.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no records affected, check the provided data")
	}

	if err != nil {
		return err
	}

	return nil
}
