package entity

import (
	"context"
	"database/sql"
)

type FlightRepository struct {
	db *sql.DB
}

func NewFlightRepository(db *sql.DB) *FlightRepository {
	return &FlightRepository{db: db}
}

func (r *FlightRepository) Save(ctx context.Context, flight Flight) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO flight VALUES (NULL,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`,
		flight.FRID,
		flight.ICAORegistration,
		flight.ICAOModel,
		flight.Latitude,
		flight.Longitude,
		flight.Heading,
		flight.Altitude,
		flight.Speed,
		flight.SquawkCode,
		flight.RadarID,
		flight.Registration,
		flight.Timestamp,
		flight.Origin,
		flight.Destination,
		flight.FlightNumber,
		flight.RateOfClimb,
		flight.CallSign,
		flight.Company,
		flight.IsOnGround,
		flight.IsGlider,
		flight.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
