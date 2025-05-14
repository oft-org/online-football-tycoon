INSERT INTO oft.team (id, name, country) VALUES
    (gen_random_uuid(), 'FC Barcelona', 'ESP'),
    (gen_random_uuid(), 'Manchester City', 'ENG');


  


INSERT INTO oft.player (id, team_id, firstname, lastname, nationality, position, age, fee, salary, technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Marc-André', 'ter Stegen', 'DEU', 'goalkeeper', 32, 35000000, 700000, 85, 88, 80, 0, false, 95, 90, 93),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Ronald', 'Araújo', 'URU', 'defender', 25, 60000000, 850000, 80, 85, 90, 0, false, 90, 92, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Jules', 'Koundé', 'FRA', 'defender', 26, 55000000, 800000, 82, 84, 85, 0, false, 88, 91, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Alejandro', 'Balde', 'ESP', 'defender', 21, 50000000, 750000, 83, 80, 85, 0, false, 85, 87, 90),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Frenkie', 'de Jong', 'NLD', 'midfielder', 27, 90000000, 1000000, 91, 89, 82, 0, false, 92, 90, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Ilkay', 'Gündogan', 'DEU', 'midfielder', 34, 30000000, 950000, 90, 88, 80, 0, false, 95, 87, 92),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Raphinha', 'Dias', 'BRA', 'forward', 28, 60000000, 900000, 88, 85, 86, 0, false, 90, 88, 93),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Robert', 'Lewandowski', 'POL', 'forward', 36, 45000000, 1100000, 92, 90, 85, 0, false, 97, 91, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'João', 'Felix', 'POR', 'midfielder', 25, 70000000, 850000, 90, 84, 82, 0, false, 89, 85, 91),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Lionel', 'Messi', 'ARG', 'forward', 36, 50000000, 1000000, 95, 90, 80, 0, false, 100, 90, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Pedri', 'Gonzalez', 'ESP', 'midfielder', 21, 70000000, 800000, 90, 85, 75, 0, false, 90, 88, 92);







INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness
)
VALUES
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Ederson', 'Moraes', 'BRA', 'goalkeeper', 31, 40000000, 800000, 84, 87, 85, 0, false, 95, 91, 93),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Kyle', 'Walker', 'ENG', 'defender', 34, 50000000, 850000, 80, 85, 92, 0, false, 90, 92, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Rúben', 'Dias', 'POR', 'defender', 27, 70000000, 950000, 83, 88, 90, 0, false, 92, 91, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'John', 'Stones', 'ENG', 'defender', 30, 60000000, 900000, 82, 87, 85, 0, false, 89, 90, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Joško', 'Gvardiol', 'CRO', 'defender', 24, 80000000, 1000000, 85, 86, 88, 0, false, 91, 90, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Rodri', 'Hernandez', 'ESP', 'midfielder', 29, 100000000, 1200000, 91, 90, 87, 0, false, 97, 93, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Bernardo', 'Silva', 'POR', 'midfielder', 30, 90000000, 1100000, 93, 89, 82, 0, false, 95, 92, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Jack', 'Grealish', 'ENG', 'forward', 29, 70000000, 900000, 89, 85, 84, 0, false, 91, 89, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Phil', 'Foden', 'ENG', 'forward', 25, 95000000, 1100000, 92, 87, 85, 0, false, 96, 93, 97),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Erling', 'Haaland', 'NOR', 'forward', 24, 150000000, 1200000, 88, 85, 95, 0, false, 95, 92, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Kevin', 'De Bruyne', 'BEL', 'midfielder', 33, 80000000, 1100000, 93, 90, 85, 0, false, 100, 90, 94);






INSERT INTO oft.strategy (id, team_id, formation, playing_style, game_tempo, passing_style, defensive_positioning, build_up_play, attack_focus, key_player_usage)
VALUES
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), '4-3-3', 'possession', 'fast_tempo', 'short', 'zonal_marking', 'play_from_back', 'wide_play', 'reference_player'),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), '4-2-3-1', 'direct', 'balanced_tempo', 'long', 'man_marking', 'counter_attack', 'central_play', 'playmaker');






INSERT INTO oft.match (id, home_team, away_team, match_date, home_result, away_result)
VALUES
    (gen_random_uuid(), 
     (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 
     (SELECT id FROM oft.team WHERE name = 'Manchester City'),
     '2025-03-15 20:00:00', 2, 2);
