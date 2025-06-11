BEGIN;

UPDATE oft.tournament AS t1
SET promotion_to = t2.id
FROM oft.tournament t2
WHERE t1.name = 'Segunda División' AND t2.name = 'Primera División';

UPDATE oft.tournament AS t1
SET descent_to = t2.id
FROM oft.tournament t2
WHERE t1.name = 'Primera División' AND t2.name = 'Segunda División';

COMMIT;
