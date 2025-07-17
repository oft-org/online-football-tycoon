UPDATE oft.match
SET
  home_result = $2,
  away_result = $3
WHERE id = $1
