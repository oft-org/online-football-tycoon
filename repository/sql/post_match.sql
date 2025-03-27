INSERT INTO eft.match (
    id,
    home_team,
    away_team,
    date,
    home_result,
    away_result
) VALUES (
    $1, $2, $3, $4, $5, $6
);
