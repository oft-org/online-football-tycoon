SELECT
    hst.formation AS home_formation,
    hst.playing_style AS home_playing_style,
    hst.game_tempo AS home_game_tempo,
    hst.passing_style AS home_passing_style,
    hst.defensive_positioning AS home_defensive_positioning,
    hst.build_up_play AS home_build_up_play,
    hst.attack_focus AS home_attack_focus,
    hst.key_player_usage AS home_key_player_usage,
    
    ast.formation AS away_formation,
    ast.playing_style AS away_playing_style,
    ast.game_tempo AS away_game_tempo,
    ast.passing_style AS away_passing_style,
    ast.defensive_positioning AS away_defensive_positioning,
    ast.build_up_play AS away_build_up_play,
    ast.attack_focus AS away_attack_focus,
    ast.key_player_usage AS away_key_player_usage
FROM oft.match m
JOIN oft.strategies hst ON hst.team_id = m.home_team
JOIN oft.strategies ast ON ast.team_id = m.away_team
WHERE m.id = $1;
