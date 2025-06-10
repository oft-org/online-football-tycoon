BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS oft.team (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    country CHAR(3) NOT NULL REFERENCES oft.country(code)
);

INSERT INTO oft.team (id, name, country) VALUES
    (gen_random_uuid(), 'Club Deportivo Bahía Real', 'ESP'),
    (gen_random_uuid(), 'Atlético Sierra Norte', 'ESP'),
    (gen_random_uuid(), 'Unión Deportiva Costa Verde', 'ESP'),
    (gen_random_uuid(), 'Fútbol Club Valle Azul', 'ESP'),
    
    (gen_random_uuid(), 'Club Atlético Rocafuerte', 'ESP'),
    (gen_random_uuid(), 'Deportivo Villa del Mar', 'ESP'),
    (gen_random_uuid(), 'Agrupación Deportiva Puente Largo', 'ESP'),
    (gen_random_uuid(), 'Sporting Monteluz CF', 'ESP');

COMMIT;
