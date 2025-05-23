SELECT
    id,
    firstname,
    lastname,
    position,
    technique,
    mental,
    physique
FROM oft.player
WHERE team_id = $1;
