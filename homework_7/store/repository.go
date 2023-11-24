package store

import (
	"context"

	"github.com/UnderMountain96/ITEA_GO/student_testing/model"
	"github.com/google/uuid"
)

type TestRepository interface {
	GetAll(ctx context.Context) ([]model.Test, error)
	GetQuestions(ctx context.Context, id uuid.UUID) ([]model.Question, error)
}
