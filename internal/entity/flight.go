package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Flight struct {
	ID               int       `bun:",pk,autoincrement"`
	FRID             string    `bun:"fr_id"`         // FlightRadar24 ID
	ICAORegistration string    `bun:"icao_reg"`      // ICAO 24-bit address The ICAO24 code (sometimes called the Mode S code) is a 24-bit unique number that is assigned to each vehicle or object that can transmit ADS-B messages. It is usually transmitted by aircraft but some airport ground vehicles and multilateration towers also have ICAO24 codes assigned to them.
	Latitude         float64   `bun:"lat"`           // Latitude
	Longitude        float64   `bun:"lon"`           // Longitude
	Heading          uint8     `bun:"heading"`       // Heading in degrees
	Altitude         uint      `bun:"alt"`           // Altitude in feet
	Speed            uint      `bun:"speed"`         // Speed in knots
	SquawkCode       string    `bun:"squawk_code"`   // Squawk code
	RadarID          string    `bun:"radar_id"`      // Radar ID
	ICAOModel        string    `bun:"icao_model"`    // ICAO model type
	Registration     string    `bun:"registration"`  // Registration
	Timestamp        int64     `bun:"timestamp"`     // Timestamp
	Origin           string    `bun:"origin"`        // Origin airport IATA code
	Destination      string    `bun:"destination"`   // Destination airport IATA code
	FlightNumber     string    `bun:"flight_number"` // Flight number
	IsOnGround       bool      `bun:"is_on_ground"`  // Is on ground
	RateOfClimb      uint      `bun:"is_glider"`     // Rate of climb in feet per minute
	CallSign         string    `bun:"call_sign"`     // Callsign
	IsGlider         bool      `bun:"is_glider"`     // Is glider
	Company          string    `bun:"company"`       // Company
	CreatedAt        time.Time `bun:"created_at"`
	bun.BaseModel    `bun:"table:flight,alias:f"`
}
