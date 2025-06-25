SELECT 
te.team_id,
te.team_name,
cl.points,
cl.goalsFor,
cl.goalsAgainst,
(cl.goalsFor-cl.goalsAgainst) AS goal_difference
FROM oft.season_team st 
JOIN oft.team te ON oft.season_team.team_id
JOIN oft.classification cl ON oft.classification.team_id = oft.team.team_id
WHERE st.season_id = $1
ORDER BY oft.classification.points DESC, goal_difference DESC;