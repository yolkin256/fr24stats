package entity

import (
	"context"
	"github.com/uptrace/bun"
)

type FlightRepository struct {
	db *bun.DB
}

func NewFlightRepository(db *bun.DB) *FlightRepository {
	return &FlightRepository{db: db}
}

func (r *FlightRepository) Save(ctx context.Context, flights []Flight) error {
	if len(flights) == 0 {
		return nil
	}

	_, err := r.db.NewInsert().Model(&flights).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
