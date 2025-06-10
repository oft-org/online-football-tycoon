BEGIN;

DELETE FROM oft.match
WHERE id IN (
  SELECT m.id
  FROM oft.match m
  JOIN oft.season s ON m.season_id = s.id
  WHERE m.home_result IS NULL AND m.away_result IS NULL
);

COMMIT;
