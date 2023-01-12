package repository

import (
	"context"

	"github.com/jsierra3991/platzi/proto/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error

	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(cxt context.Context, test *models.Test) error
	SetQuestion(cxt context.Context, question *models.Question) error

	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error)
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return impl.SetStudent(ctx, student)
}
func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return impl.GetStudent(ctx, id)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return impl.GetTest(ctx, id)
}
func SetTest(ctx context.Context, test *models.Test) error {
	return impl.SetTest(ctx, test)
}
func SetQuestion(ctx context.Context, question *models.Question) error {
	return impl.SetQuestion(ctx, question)
}

func SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return impl.SetEnrollment(ctx, enrollment)
}
func GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return impl.GetStudentPerTest(ctx, testId)
}

func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return impl.GetQuestionsPerTest(ctx, testId)
}
