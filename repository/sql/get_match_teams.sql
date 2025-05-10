SELECT
    ht.id AS home_team_id,
    ht.name AS home_team_name,
    at.id AS away_team_id,
    at.name AS away_team_name
FROM oft.match m
JOIN oft.teams ht ON m.home_team = ht.id
JOIN oft.teams at ON m.away_team = at.id
WHERE m.id = $1;
