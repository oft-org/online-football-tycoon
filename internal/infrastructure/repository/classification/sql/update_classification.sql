UPDATE oft.classification
SET
  points = $2,
  goals_for = $3,
  goals_against = $4
WHERE team_id = $1
