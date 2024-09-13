package input

import (
	"context"
	"time"
)

type Input struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]Input, error)
}

type StaticInputs struct {
}

func (s StaticInputs) GetAll(ctx context.Context) ([]Input, error) {
	return []Input{
		{
			ID:        1,
			Name:      "Option A",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Option B",
			CreatedAt: time.Now().Add(time.Hour * 24 * -2),
		},
		{
			ID:        3,
			Name:      "Option B",
			CreatedAt: time.Now().Add(time.Hour * 24 * -3),
		},
	}, nil
}
