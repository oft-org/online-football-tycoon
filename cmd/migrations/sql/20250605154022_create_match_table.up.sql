BEGIN;

CREATE TABLE IF NOT EXISTS oft.match (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    season_id UUID NOT NULL REFERENCES oft.season(id) ON DELETE CASCADE,
    home_team uuid REFERENCES oft.team(id) ON DELETE CASCADE,
    away_team uuid REFERENCES oft.team(id) ON DELETE CASCADE,
    match_date TIMESTAMP,
    home_result INT,
    away_result INT
   );

COMMIT;
