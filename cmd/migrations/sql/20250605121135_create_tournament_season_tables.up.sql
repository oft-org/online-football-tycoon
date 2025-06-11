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

COMMIT;
