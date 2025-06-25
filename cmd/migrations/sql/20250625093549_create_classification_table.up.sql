BEGIN;

CREATE TABLE IF NOT EXISTS oft.classification (
    team_id UUID PRIMARY KEY DEFAULT,
    points INT,
    goalsFor INT,
    goalsAgainst INT
);

COMMIT;
