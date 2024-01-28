-- +goose Up
CREATE TABLE flight
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    fr_id         TEXT,
    icao_reg      TEXT,
    icao_model    TEXT,
    lat           FLOAT,
    lon           FLOAT,
    heading       INTEGER,
    alt           INT,
    speed         INT,
    squawk_code   TEXT,
    radar_id      TEXT,
    registration  TEXT,
    timestamp     INTEGER,
    origin        TEXT,
    destination   TEXT,
    flight_number TEXT,
    rate_of_climb INTEGER,
    call_sign     INTEGER,
    company       TEXT,
    is_on_ground  INTEGER,
    is_glider     INTEGER,
    created_at    DATETIME
);
