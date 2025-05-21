CREATE TABLE IF NOT EXISTS oft.match_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id UUID NOT NULL REFERENCES oft.match(id) ON DELETE CASCADE,
    team_id UUID NOT NULL REFERENCES oft.team(id) ON DELETE CASCADE,
    event_type VARCHAR(255) NOT NULL,
    minute INT CHECK (minute >= 0),
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
