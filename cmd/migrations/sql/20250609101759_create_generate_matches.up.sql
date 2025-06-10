BEGIN;

INSERT INTO oft.match (season_id, home_team, away_team)
SELECT
  s1.season_id,
  s1.team_id AS home_team,
  s2.team_id AS away_team
FROM oft.season_team s1
JOIN oft.season_team s2
  ON s1.season_id = s2.season_id
WHERE s1.team_id <> s2.team_id;

COMMIT;
