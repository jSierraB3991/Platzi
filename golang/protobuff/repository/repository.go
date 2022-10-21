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
