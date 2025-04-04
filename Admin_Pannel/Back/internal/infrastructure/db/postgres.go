// internal/infrastructure/db/postgres.go
package db

import (
	"database/sql"
	"ticket-service/internal/entity"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetPosts() ([]entity.Post, error) {
	query := `SELECT id, user_id, name, age, tickets, data, status FROM post`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var p entity.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Age, &p.Tickets, &p.Data, &p.Status)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostgresRepo) GetPrePosts() ([]entity.PrePost, error) {
	query := `SELECT id, user_id, name, age, tickets FROM pre_post`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prePosts []entity.PrePost
	for rows.Next() {
		var p entity.PrePost
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Age, &p.Tickets)
		if err != nil {
			return nil, err
		}
		prePosts = append(prePosts, p)
	}

	return prePosts, nil
}
