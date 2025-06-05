BEGIN;

CREATE TABLE IF NOT EXISTS oft.tournament (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('League', 'Cup')),
    country_code CHAR(3) NOT NULL REFERENCES oft.country(code),
    division INT NOT NULL CHECK (division >= 1),
    promotion_to UUID REFERENCES oft.tournament(id) ON DELETE SET NULL,
    descent_to UUID REFERENCES oft.tournament(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS oft.season (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id UUID NOT NULL REFERENCES oft.tournament(id) ON DELETE CASCADE,
    from_date DATE NOT NULL,
    to_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS oft.season_team (
    season_id UUID NOT NULL REFERENCES oft.season(id) ON DELETE CASCADE,
    team_id UUID NOT NULL REFERENCES oft.team(id) ON DELETE CASCADE,
    PRIMARY KEY (season_id, team_id)
);

WITH primera AS (
    INSERT INTO oft.tournament (id, name, type, country_code, division, promotion_to, descent_to)
    VALUES (gen_random_uuid(), 'Primera División', 'League', 'ESP', 1, NULL, NULL)
    RETURNING id
),
segunda AS (
    INSERT INTO oft.tournament (id, name, type, country_code, division, promotion_to, descent_to)
    SELECT gen_random_uuid(), 'Segunda División', 'League', 'ESP', 2, primera.id, NULL
    FROM primera
),
update_primera_descent AS (
    UPDATE oft.tournament
    SET descent_to = (SELECT id FROM segunda)
    WHERE id = (SELECT id FROM primera)
),
cup AS (
    INSERT INTO oft.tournament (id, name, type, country_code, division, promotion_to, descent_to)
    VALUES (gen_random_uuid(), 'Copa de España', 'Cup', 'ESP', 1, NULL, NULL)
    RETURNING id
),
tournaments AS (
    SELECT id, name FROM oft.tournament WHERE name IN ('Primera División', 'Segunda División', 'Copa de España')
),
insert_seasons AS (
    INSERT INTO oft.season (id, tournament_id, from_date, to_date)
    SELECT
        gen_random_uuid(),
        id,
        DATE '2025-08-01',
        DATE '2026-05-30'
    FROM tournaments
    RETURNING id, tournament_id
),
seasons AS (
    SELECT s.id as season_id, t.name as tournament_name
    FROM oft.season s
    JOIN oft.tournament t ON s.tournament_id = t.id
    WHERE t.name IN ('Primera División', 'Segunda División', 'Copa de España')
),
teams AS (
    SELECT id as team_id FROM oft.team LIMIT 5
),
insert_season_teams AS (
    INSERT INTO oft.season_team (season_id, team_id)
    SELECT season_id, team_id
    FROM seasons CROSS JOIN teams
    RETURNING season_id, team_id
)
SELECT * FROM insert_season_teams;

COMMIT;
