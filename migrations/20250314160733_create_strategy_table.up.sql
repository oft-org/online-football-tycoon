CREATE TABLE IF NOT EXISTS oft.strategy (
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

INSERT INTO oft.strategy (id, team_id, formation, playing_style, game_tempo, passing_style, defensive_positioning, build_up_play, attack_focus, key_player_usage)
VALUES
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), '4-3-3', 'possession', 'fast_tempo', 'short', 'zonal_marking', 'play_from_back', 'wide_play', 'reference_player'),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), '4-2-3-1', 'direct', 'balanced_tempo', 'long', 'man_marking', 'counter_attack', 'central_play', 'playmaker');
