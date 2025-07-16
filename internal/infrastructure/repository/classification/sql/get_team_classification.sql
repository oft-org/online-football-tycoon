SELECT
  points,
  goals_for,
  goals_against
FROM oft.classification
WHERE team_id = $1
