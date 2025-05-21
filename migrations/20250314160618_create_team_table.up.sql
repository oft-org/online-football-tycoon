CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS oft.team (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    country CHAR(3) NOT NULL REFERENCES oft.country(code)
);

INSERT INTO oft.team (id, name, country) VALUES
    (gen_random_uuid(), 'FC Barcelona', 'ESP'),
    (gen_random_uuid(), 'Manchester City', 'GBR');
