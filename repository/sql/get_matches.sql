SELECT 
 m.id AS match_id,
    ht.name AS home_team_name,
    at.name AS away_team_name,
    hs.formation AS home_formation,
    hs.playing_style AS home_playing_style,
    hs.game_tempo AS home_game_tempo,
    hs.passing_style AS home_passing_style,
    hs.defensive_positioning AS home_defensive_positioning,
    hs.build_up_play AS home_build_up_play,
    hs.attack_focus AS home_attack_focus,
    hs.key_player_usage AS home_key_player_usage,
    away_strategy.formation AS away_formation,
    away_strategy.playing_style AS away_playing_style,
    away_strategy.game_tempo AS away_game_tempo,
    away_strategy.passing_style AS away_passing_style,
    away_strategy.defensive_positioning AS away_defensive_positioning,
    away_strategy.build_up_play AS away_build_up_play,
    away_strategy.attack_focus AS away_attack_focus,
    away_strategy.key_player_usage AS away_key_player_usage
FROM oft.match m
JOIN oft.teams ht ON m.home_team = ht.id
JOIN oft.teams at ON m.away_team = at.id
JOIN oft.strategies hs ON hs.team_id = m.home_team
JOIN oft.strategies away_strategy ON away_strategy.team_id = m.away_team
ORDER BY m.match_date ASC;
