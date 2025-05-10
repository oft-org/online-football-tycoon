SELECT
    hp.firstname AS home_player_firstname,
    hp.lastname AS home_player_lastname,
    hp.position AS home_player_position,
    hp.technique AS home_player_technique,
    hp.mental AS home_player_mental,
    hp.physique AS home_player_physique,

    ap.firstname AS away_player_firstname,
    ap.lastname AS away_player_lastname,
    ap.position AS away_player_position,
    ap.technique AS away_player_technique,
    ap.mental AS away_player_mental,
    ap.physique AS away_player_physique
FROM oft.match m
JOIN oft.players hp ON hp.id = m.home_team
JOIN oft.players ap ON ap.id = m.away_team
WHERE m.id = $1;