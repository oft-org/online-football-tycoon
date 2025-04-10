INSERT INTO oft.match_events (
    match_id,
    team_id,
    event_type,
    minute,
    description
) VALUES (
    $1, $2, $3, $4, $5
);
