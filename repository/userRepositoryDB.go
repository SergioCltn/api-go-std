package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/sergiocltn/api-go-std/models"
)

type IUserRepository interface {
	CreateUserQuery(user models.User) (int, error)
	GetUserByIDQuery(id int) (*models.User, error)
	ListUsersQuery() ([]models.User, error)
	UpdateUserQuery(id int, user models.User) error
	DeleteUserQuery(id int) error
}

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepositoryDB(connStr string) *UserRepositoryDB {
	db, err := sql.Open("postgres", connStr)
	fmt.Println("DB connection established")
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return nil
	}
	return &UserRepositoryDB{
		DB: db,
	}
}

func (m *UserRepositoryDB) Close(user models.User) {
	m.DB.Close()
	fmt.Println("DB has been closed")
}

func (m *UserRepositoryDB) CreateUserQuery(user models.User) (int, error) {
	sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := m.DB.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoModels
		} else {
			return 0, err
		}
	}
	return id, nil
}

func (m *UserRepositoryDB) GetUserByIDQuery(id int) (*models.User, error) {
	sqlStatement := `SELECT id, name, email, password FROM users WHERE id=$1`
	row := m.DB.QueryRow(sqlStatement, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.User{}, models.ErrNoModels
		} else {
			return &models.User{}, err
		}
	}
	return &user, nil
}

func (m *UserRepositoryDB) ListUsersQuery() ([]models.User, error) {
	sqlStatement := `SELECT id, name, email, password FROM users`
	rows, err := m.DB.Query(sqlStatement)
	if err != nil {
		slog.Error("Failed to list users", "error", err)
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			slog.Error("Failed to scan user row", "error", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		slog.Error("Failed to iterate over user rows", "error", err)
		return nil, err
	}
	return users, nil
}

func (m *UserRepositoryDB) UpdateUserQuery(id int, user models.User) error {
	sqlStatement := `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`
	_, err := m.DB.Exec(sqlStatement, user.Name, user.Email, user.Password, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoModels
		} else {
			return err
		}
	}
	return nil
}

func (m *UserRepositoryDB) DeleteUserQuery(id int) error {
	sqlStatement := `DELETE FROM users WHERE id=$1`
	_, err := m.DB.Exec(sqlStatement, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrNoModels
		} else {
			return err
		}
	}
	return nil
}
