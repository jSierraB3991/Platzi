package database

import (
	"context"
	"database/sql"

	"github.com/jsierra3991/platzi/proto/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgreRepository(url string) *PostgresRepository {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO student (id, name, age) VALUES ($1, $2, $3)",
		student.Id, student.Name, student.Age)
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM student WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			return err
		}
	}()

	var student = models.Student{}

	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return err
		}
	}
	return &student, nil
}
