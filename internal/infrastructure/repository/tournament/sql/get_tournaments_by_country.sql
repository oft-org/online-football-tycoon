SELECT 
    t.id,
    t.name,
    t.type,
    t.country_code,
    t.division,
    t.promotion_to,
    t.descent_to
FROM oft.tournament t
WHERE t.country_code = $1;
