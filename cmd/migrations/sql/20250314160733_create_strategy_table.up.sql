BEGIN;

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
    
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), '4-3-3', 'possession', 'fast_tempo', 'short', 'zonal_marking', 'play_from_back', 'wide_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), '3-4-3', 'high_press', 'balanced_tempo', 'short', 'man_marking', 'play_from_back', 'central_play', 'free_role_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), '4-3-3', 'counter_attack', 'fast_tempo', 'long', 'zonal_marking', 'long_clearance', 'central_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), '5-3-2', 'direct_play', 'slow_tempo', 'long', 'man_marking', 'long_clearance', 'wide_play', 'free_role_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), '4-4-2', 'low_block', 'slow_tempo', 'short', 'zonal_marking', 'play_from_back', 'central_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), '3-4-3', 'possession', 'balanced_tempo', 'short', 'zonal_marking', 'play_from_back', 'wide_play', 'free_role_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), '3-4-3', 'high_press', 'fast_tempo', 'short', 'man_marking', 'play_from_back', 'wide_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Sporting Monteluz CF'), '4-4-2', 'direct_play', 'balanced_tempo', 'long', 'man_marking', 'long_clearance', 'central_play', 'free_role_player');

COMMIT;
