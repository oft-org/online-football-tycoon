SELECT
    formation,
    playing_style,
    game_tempo,
    passing_style,
    defensive_positioning,
    build_up_play,
    attack_focus,
    key_player_usage
FROM oft.strategies
WHERE team_id = $1;
