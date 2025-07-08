BEGIN;

CREATE TABLE IF NOT EXISTS oft.classification (
    team_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    points INT,
    goals_for INT,
    goals_against INT
);

COMMIT;
