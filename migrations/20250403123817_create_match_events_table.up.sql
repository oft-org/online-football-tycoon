BEGIN;
CREATE TABLE oft.match_events (
    match_id UUID NOT NULL REFERENCES oft.match(id) ON DELETE CASCADE,
    home_events JSONB NOT NULL,
    away_events JSONB NOT NULL,
    home_score_chances INT CHECK (home_score_chances >= 0),
    away_score_chances INT CHECK (away_score_chances >= 0),
    home_goals INT CHECK (home_goals >= 0),
    away_goals INT CHECK (away_goals >= 0),
    created_at TIMESTAMP DEFAULT NOW()
);
COMMIT;
