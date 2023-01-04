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

func NewPostgreRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO student (id, name, age) VALUES ($1, $2, $3)",
		student.Id, student.Name, student.Age)
	return err
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions(id, answer, question, test_id) VALUES ($1, $2, $3, $4)",
		question.Id, question.Answer, question.Question, question.TestId)
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM student WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
	}()

	var student = models.Student{}

	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
	}
	return &student, nil
}

func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM test WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
	}()

	var test = models.Test{}

	for rows.Next() {
		err = rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
	}
	return &test, nil
}
func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO test (id, name) VALUES ($1, $2)",
		test.Id, test.Name)
	return err

}
