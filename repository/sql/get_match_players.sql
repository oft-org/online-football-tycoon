SELECT
    id,
    firstname,
    lastname,
    position,
    technique,
    mental,
    physique
FROM oft.players
WHERE team_id = $1;