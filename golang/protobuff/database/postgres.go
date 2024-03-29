package database

import (
	"context"
	"database/sql"
	"log"

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
func (repo *PostgresRepository) SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO enrollment (student_id, test_id) VALUES ($1, $2)",
		enrollment.StudentId, enrollment.TestId)
	return err
}
func (repo *PostgresRepository) GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT s.id, s.name, s.age FROM student s INNER JOIN enrollment e ON s.id = e.student_id WHERE e.test_id = $1", testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var students = []*models.Student{}

	for rows.Next() {
		var student = models.Student{}
		if err = rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, &student)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, rows.Err()
	}
	return students, nil
}

func (repo *PostgresRepository) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, question FROM questions WHERE test_id = $1", testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var questions = []*models.Question{}
	for rows.Next() {
		var question = models.Question{}
		if err = rows.Scan(&question.Id, &question.Question); err == nil {
			questions = append(questions, &question)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, rows.Err()
	}
	return questions, nil
}
