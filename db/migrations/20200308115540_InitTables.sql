
-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE athletes(
    chip_id UUID PRIMARY KEY,
    start_number INTEGER NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    CONSTRAINT start_number_uniq UNIQUE (start_number)
);

CREATE TABLE timing_points(
    timing_id UUID PRIMARY KEY,
    point_id VARCHAR(20) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    chip_id UUID REFERENCES athletes(chip_id) NOT NULL
);

-- +goose Down
DROP EXTENSION IF EXISTS "uuid-ossp";

DROP TABLE timing_points;

DROP TABLE athletes;

