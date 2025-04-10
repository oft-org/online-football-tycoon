INSERT INTO oft.match_events (
    match_id,
    home_events,
    away_events,
    home_score_chances,
    away_score_chances,
    home_goals,
    away_goals
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);
