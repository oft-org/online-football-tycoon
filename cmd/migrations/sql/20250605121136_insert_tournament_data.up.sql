BEGIN;

INSERT INTO oft.tournament (id, name, type, country_code, division)
VALUES 
  (gen_random_uuid(), 'Primera División', 'League', 'ESP', 1),
  (gen_random_uuid(), 'Segunda División', 'League', 'ESP', 2),
  (gen_random_uuid(), 'Copa de España', 'Cup', 'ESP', 1);

COMMIT;
