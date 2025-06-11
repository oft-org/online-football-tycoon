BEGIN;

INSERT INTO oft.season (id, tournament_id, from_date, to_date)
SELECT gen_random_uuid(), t.id, DATE '2025-08-01', DATE '2026-05-30'
FROM oft.tournament t
WHERE t.name IN ('Primera División', 'Segunda División', 'Copa de España')
  AND NOT EXISTS (
    SELECT 1 FROM oft.season s WHERE s.tournament_id = t.id
  );

INSERT INTO oft.season_team (season_id, team_id)
SELECT s.id, t.id
FROM oft.season s
JOIN oft.tournament tor ON s.tournament_id = tor.id
JOIN oft.team t ON (
  (tor.name = 'Primera División' AND t.name IN (
    'Club Deportivo Bahía Real',
    'Atlético Sierra Norte',
    'Unión Deportiva Costa Verde',
    'Fútbol Club Valle Azul'
  )) OR
  (tor.name = 'Segunda División' AND t.name IN (
    'Club Atlético Rocafuerte',
    'Deportivo Villa del Mar',
    'Agrupación Deportiva Puente',
    'Sporting Monteluz CF'
  )) OR
  (tor.name = 'Copa de España' AND t.name IN (
    'Club Deportivo Bahía Real',
    'Atlético Sierra Norte',
    'Unión Deportiva Costa Verde',
    'Fútbol Club Valle Azul',
    'Club Atlético Rocafuerte',
    'Deportivo Villa del Mar',
    'Agrupación Deportiva Puente',
    'Sporting Monteluz CF'
  ))
);

COMMIT;
