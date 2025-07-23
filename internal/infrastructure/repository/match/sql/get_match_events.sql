SELECT
    id,
    match_id,
    team_id,
    event_type,
    minute,
    description
FROM oft.match_events
WHERE match_id = $1
ORDER BY created_at ASC;
