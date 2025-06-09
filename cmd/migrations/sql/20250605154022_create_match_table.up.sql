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

INSERT INTO oft.match (id, season_id, home_team, away_team, match_date, home_result, away_result)
VALUES (
    gen_random_uuid(),
    (SELECT id FROM oft.season WHERE tournament_id = (SELECT id FROM oft.tournament WHERE name = 'Primera División') LIMIT 1),
    (SELECT id FROM oft.team WHERE name = 'FC Barcelona'),
    (SELECT id FROM oft.team WHERE name = 'Manchester City'),
    '2025-03-15 20:00:00',
    2,
    2
);

COMMIT;
