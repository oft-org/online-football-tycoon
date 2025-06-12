INSERT INTO oft.match (
    season_id,
    home_team,
    away_team,
    match_date,
    home_result,
    away_result
) VALUES (
    $1, $2, $3, $4, $5, $6
);
