package repository

import (
	"context"

	"github.com/jsierra3991/platzi/proto/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func SetStudent(ctx context.Context, student *models.Student) {
	return impl.SetSetudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) {
	return impl.GetStudent(ctx, id)
}
