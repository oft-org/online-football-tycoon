BEGIN;

CREATE TABLE oft.strategy (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id UUID REFERENCES oft.team(id) ON DELETE CASCADE,
    formation VARCHAR(255),
    playing_style VARCHAR(255),
    game_tempo VARCHAR(255),
    passing_style VARCHAR(255),
    defensive_positioning VARCHAR(255),
    build_up_play VARCHAR(255),
    attack_focus VARCHAR(255),
    key_player_usage VARCHAR(255)
);

COMMIT;
