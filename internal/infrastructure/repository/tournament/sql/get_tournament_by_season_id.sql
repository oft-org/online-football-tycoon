SELECT 
    t.id,
    t.name,
    t.type,
    t.country_code,
    t.division,
    t.promotion_to,
    t.descent_to
FROM oft.season s
JOIN oft.tournament t ON s.tournament_id = t.id
WHERE s.id = $1;
