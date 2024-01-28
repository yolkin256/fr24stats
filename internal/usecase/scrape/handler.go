package scrape

import (
	"context"
	"fr24stats/internal/entity"
	"github.com/genericplatform/flightradar24sdk"
	"log"
	"time"
)

type Cmd struct {
	Airline string
}

type Handler struct {
	api        *flightradar24sdk.API
	repository *entity.FlightRepository
}

func NewHandler(api *flightradar24sdk.API, repository *entity.FlightRepository) *Handler {
	return &Handler{api: api, repository: repository}
}

func (h *Handler) Handle(ctx context.Context, cmd Cmd) error {
	res, err := h.api.GetFlights(ctx, cmd.Airline, nil)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Получено %d записей о полётах для авиалинии %s\n", len(res.Flights), cmd.Airline)

	for _, f := range res.Flights {
		if err := h.repository.Save(ctx, entity.Flight{
			FRID:             f.ID,
			ICAORegistration: f.ICAORegistration,
			Latitude:         f.Latitude,
			Longitude:        f.Longitude,
			Heading:          f.Heading,
			Altitude:         f.Altitude,
			Speed:            f.Speed,
			SquawkCode:       f.SquawkCode,
			RadarID:          f.RadarID,
			ICAOModel:        f.ICAOModel,
			Registration:     f.Registration,
			Timestamp:        f.Timestamp,
			Origin:           f.Origin,
			Destination:      f.Destination,
			FlightNumber:     f.FlightNumber,
			IsOnGround:       f.IsOnGround,
			RateOfClimb:      f.RateOfClimb,
			CallSign:         f.CallSign,
			IsGlider:         f.IsGlider,
			Company:          f.Company,
			CreatedAt:        time.Now(),
		}); err != nil {
			return err
		}
	}

	return nil
}
