package repository

import (
	"database/sql"
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
func (repository Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.db.Exec(
		`INSERT INTO Users (name, email,email_hash, safepassword) VALUES (?,?, ?,?)`,
		user.Name,
		user.Email,
		user.Email_Hash,
		user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no records affected, check the provided data")
		} else {
			return 0, err
		}
	}

	userID, err := statement.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(userID), nil
}

// FindByID finds a user in the database by ID.
func (repository Users) FindByID(userID uint64) (models.User, error) {
	user := models.User{}
	err := repository.db.Get(&user,
		`SELECT id, name, email, created_at FROM Users WHERE id = ?`,
		userID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("no records affected, check the provided data")
		} else {
			return models.User{}, err
		}
	}
	return user, nil
}

// FindAllUsers retrieves all users saved in the database.
func (repository Users) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := repository.db.Select(&users, "SELECT id, name, email,created_at FROM Users")
	if len(users) == 0 {
		return []models.User{}, errors.New("no users found, check the provided data")
	}

	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// UpdateUser updates user information in the database.
func (repository Users) UpdateUser(userID uint64, user models.User) error {
	_, err := repository.db.Exec(
		`UPDATE Users SET name=?, email=?, safepassword=? WHERE id=?`,
		user.Name,
		user.Email,
		user.Password,
		userID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no records affected, check the provided data")
		} else {
			return err
		}
	}
	return nil
}

// FindByEmail finds a user by email_hash and returns its ID and hashed safepassword.
func (repository Users) FindByEmail(email_hash string) (models.User, error) {
	user := models.User{}
	err := repository.db.Get(&user, "SELECT id, safepassword FROM Users WHERE email_hash = ?", email_hash)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("no records affected, check the provided data")
		} else {
			return models.User{}, err
		}
	}
	return user, nil
}

// DeleteUser deletes a user from the database.
func (repository Users) DeleteUser(userID uint64) error {
	_, err := repository.db.Exec(
		`DELETE FROM Users WHERE id = ?`,
		userID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no records affected, check the provided data")
		} else {
			return err
		}
	}
	return nil
}
