BEGIN;

CREATE TABLE oft.team (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    country CHAR(3) NOT NULL REFERENCES oft.country(code)
);

COMMIT;
