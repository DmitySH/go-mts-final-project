-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auto
(
    id   bigserial primary key,
    name varchar(128)
);

CREATE TABLE IF NOT EXISTS driver
(
    id   uuid primary key,
    name varchar(128),
    auto bigint references auto (id),
    lat  double precision,
    lng  double precision
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS driver;
DROP TABLE IF EXISTS auto;

-- +goose StatementEnd
