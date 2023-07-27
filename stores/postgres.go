package stores

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/fredrikaverpil/go-api-std/lib"
	"github.com/fredrikaverpil/go-api-std/models"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (PostgresStore, error) {
	driver := "postgres"
	dataSourceName := "postgres://root:secret@localhost:5432/test?sslmode=disable"
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return PostgresStore{}, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return PostgresStore{}, err
	}

	return PostgresStore{db: db}, nil
}

func (s *PostgresStore) CreateUser(username string, password string) (models.User, error) {
	hashedPassword, err := lib.HashPassword(password)
	if err != nil {
		return models.User{}, errors.New("could not hash password")
	}

	query := `
    INSERT INTO users (username, password)
    VALUES ($1, $2)
    RETURNING id, username
  `

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, username, hashedPassword)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func (s *PostgresStore) GetUserByUsername(username string) (models.User, error) {
	query := `
  SELECT id, username
  FROM users
  WHERE username = $1
  `

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, username)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (s *PostgresStore) GetUser(id int) (models.User, error) {
	query := `
  SELECT id, username
  FROM users
  WHERE id = $1
  `

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return user, nil
}
