SELECT 
RANK() OVER (ORDER BY cl.points DESC, (cl.goals_for - cl.goals_against) DESC) AS position,
te.id,
te.name,
cl.points,
cl.goals_for,
cl.goals_against,
(cl.goals_for-cl.goals_against) AS goal_difference
FROM oft.season_team st
JOIN oft.team te ON st.team_id = te.id
JOIN oft.classification cl ON cl.team_id = te.id
WHERE st.season_id = $1
ORDER BY position;