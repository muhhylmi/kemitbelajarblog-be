package repository

import (
	"database/sql"
	"fmt"
	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/model"
)

type postgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository creates a new PostgreSQL-backed user repository.
func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) FindByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, name, avatar, role, username, password_hash, created_at, updated_at
		FROM authors
		WHERE username = $1
	`

	var user model.User
	var uName, pwdHash sql.NullString
	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Name, &user.Avatar, &user.Role, &uName, &pwdHash, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error querying user by username: %w", err)
	}

	if uName.Valid {
		user.Username = uName.String
	}
	if pwdHash.Valid {
		user.PasswordHash = pwdHash.String
	}

	return &user, nil
}

func (r *postgresUserRepository) FindByID(id string) (*model.User, error) {
	query := `
		SELECT id, name, avatar, role, username, password_hash, created_at, updated_at
		FROM authors
		WHERE id = $1
	`

	var user model.User
	var uName, pwdHash sql.NullString
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.Avatar, &user.Role, &uName, &pwdHash, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error querying user by id: %w", err)
	}

	if uName.Valid {
		user.Username = uName.String
	}
	if pwdHash.Valid {
		user.PasswordHash = pwdHash.String
	}

	return &user, nil
}

func (r *postgresUserRepository) FindAll() ([]model.User, error) {
	query := `
		SELECT id, name, avatar, role, username, created_at, updated_at
		FROM authors
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying all users: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		var uName sql.NullString
		if err := rows.Scan(&user.ID, &user.Name, &user.Avatar, &user.Role, &uName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		if uName.Valid {
			user.Username = uName.String
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *postgresUserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO authors (name, avatar, role, username, password_hash)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		user.Name, user.Avatar, user.Role, user.Username, user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (r *postgresUserRepository) Update(user *model.User) error {
	// Only update name, avatar, role, username (and password if not empty, but we'll assume separate logic or skip for now)
	// Actually, if we want to allow editing, let's just update the basics.
	query := `
		UPDATE authors
		SET name = $1, avatar = $2, role = $3, username = $4, updated_at = NOW()
		WHERE id = $5
	`

	_, err := r.db.Exec(query, user.Name, user.Avatar, user.Role, user.Username, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (r *postgresUserRepository) Delete(id string) error {
	query := `DELETE FROM authors WHERE id = $1`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
