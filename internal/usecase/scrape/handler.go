package scrape

import (
	"context"
	"fr24stats/internal/entity"
	"fr24stats/internal/util"
	"github.com/genericplatform/flightradar24sdk"
	"log"
	"time"
)

const (
	insertChunkSize = 5
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

	flights := make([]entity.Flight, 0, len(res.Flights))
	now := time.Now()
	for _, f := range res.Flights {
		flights = append(flights, entity.Flight{
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
			CreatedAt:        now,
		})
	}

	log.Printf("[INFO] Получено %d записей о полётах для авиалинии %s\n", len(flights), cmd.Airline)

	chunks := util.SliceChunk(flights, insertChunkSize)
	for _, chunk := range chunks {
		if err := h.repository.Save(ctx, chunk); err != nil {
			log.Printf("[ERROR] Ошибка сохранения данных о полётах: %v", err)
			continue
		}
	}

	return nil
}
