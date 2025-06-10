INSERT INTO oft.team (id, name, country) VALUES
    (gen_random_uuid(), 'Club Deportivo Bahía Real', 'ESP'),
    (gen_random_uuid(), 'Atlético Sierra Norte', 'ESP'),
    (gen_random_uuid(), 'Unión Deportiva Costa Verde', 'ESP'),
    (gen_random_uuid(), 'Fútbol Club Valle Azul', 'ESP'),
    
    (gen_random_uuid(), 'Club Atlético Rocafuerte', 'ESP'),
    (gen_random_uuid(), 'Deportivo Villa del Mar', 'ESP'),
    (gen_random_uuid(), 'Agrupación Deportiva Puente Largo', 'ESP'),
    (gen_random_uuid(), 'Sporting Monteluz CF', 'ESP');


  


INSERT INTO oft.player (id, team_id, firstname, lastname, nationality, position, age, fee, salary, technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Álvaro', 'Jiménez', 'ESP', 'goalkeeper', 30, 7000000, 450000, 80, 82, 78, 0, false, 88, 85, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Sergio', 'Martínez', 'ESP', 'defender', 27, 8500000, 460000, 82, 80, 81, 0, false, 85, 83, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Tobias', 'Müller', 'DEU', 'defender', 28, 10000000, 520000, 83, 84, 82, 0, false, 88, 86, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Iván', 'Domingo', 'ESP', 'defender', 26, 8000000, 470000, 81, 79, 80, 0, false, 84, 83, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Lucas', 'Dubois', 'FRA', 'defender', 27, 12000000, 550000, 85, 86, 83, 0, false, 89, 87, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Matteo', 'Conti', 'ITA', 'midfielder', 30, 15000000, 600000, 88, 87, 84, 0, false, 92, 89, 91),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Leandro', 'Silva', 'BRA', 'midfielder', 25, 14000000, 570000, 86, 85, 83, 0, false, 91, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Jordi', 'Navarro', 'ESP', 'midfielder', 28, 9000000, 480000, 85, 84, 79, 0, false, 87, 84, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Raúl', 'Benítez', 'ESP', 'forward', 29, 10000000, 500000, 88, 83, 82, 0, false, 90, 86, 91),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Nicolás', 'Fernández', 'ARG', 'forward', 27, 13000000, 560000, 87, 86, 84, 0, false, 90, 88, 92),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), 'Andrej', 'Novak', 'SVK', 'forward', 29, 11000000, 530000, 85, 84, 83, 0, false, 89, 85, 90);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Álvaro', 'Romero', 'ESP', 'goalkeeper', 28, 9000000, 250000, 78, 79, 76, 0, false, 85, 87, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Sergio', 'Gálvez', 'ESP', 'defender', 27, 8500000, 230000, 77, 78, 80, 0, false, 83, 86, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Luca', 'Bianchi', 'ITA', 'defender', 30, 12000000, 280000, 82, 84, 83, 0, false, 90, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Julien', 'Moreau', 'FRA', 'defender', 28, 11500000, 275000, 80, 83, 82, 0, false, 89, 87, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Iván', 'Herrera', 'ESP', 'midfielder', 26, 9500000, 260000, 80, 82, 78, 0, false, 88, 89, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Carlos', 'Navarro', 'ESP', 'midfielder', 25, 9100000, 240000, 79, 80, 77, 0, false, 86, 88, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Tomás', 'Ferreira', 'PRT', 'midfielder', 27, 9800000, 260000, 81, 82, 79, 0, false, 87, 86, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Daniel', 'Müller', 'DEU', 'midfielder', 29, 10200000, 265000, 82, 85, 80, 0, false, 89, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Lucas', 'Schmid', 'CHE', 'forward', 30, 11000000, 290000, 83, 84, 81, 0, false, 88, 87, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Miguel', 'Luna', 'ESP', 'forward', 29, 10000000, 270000, 81, 79, 80, 0, false, 87, 86, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), 'Andre', 'Johnson', 'USA', 'forward', 26, 10800000, 285000, 84, 83, 84, 0, false, 90, 89, 91);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Javier', 'Moreno', 'ESP', 'goalkeeper', 30, 9500000, 240000, 79, 81, 78, 0, false, 85, 86, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Pablo', 'Santos', 'ESP', 'defender', 27, 8700000, 230000, 77, 78, 80, 0, false, 84, 85, 86),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Raúl', 'Castillo', 'ESP', 'defender', 28, 9000000, 235000, 80, 82, 79, 0, false, 86, 87, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Maxime', 'Dupont', 'FRA', 'defender', 29, 11000000, 270000, 82, 83, 81, 0, false, 89, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Marco', 'Rossi', 'ITA', 'defender', 30, 11500000, 280000, 83, 84, 82, 0, false, 90, 89, 91),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Niklas', 'Schneider', 'DEU', 'midfielder', 27, 10500000, 260000, 81, 82, 80, 0, false, 87, 86, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Lucas', 'Martins', 'BRA', 'midfielder', 28, 10800000, 265000, 82, 83, 81, 0, false, 88, 87, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Thiago', 'Costa', 'BRA', 'midfielder', 26, 11500000, 275000, 84, 85, 82, 0, false, 89, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'André', 'Lopes', 'PRT', 'forward', 27, 11000000, 270000, 83, 84, 81, 0, false, 88, 87, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Sergio', 'Gómez', 'ESP', 'forward', 26, 9200000, 245000, 81, 80, 77, 0, false, 87, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), 'Johan', 'Larsen', 'DNK', 'forward', 25, 10800000, 265000, 82, 83, 80, 0, false, 87, 86, 88);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Diego', 'Fernández', 'ESP', 'goalkeeper', 29, 10500000, 250000, 80, 83, 79, 0, false, 87, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Sergio', 'Navarro', 'ESP', 'defender', 28, 10000000, 245000, 78, 80, 81, 0, false, 86, 87, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Álvaro', 'Gómez', 'ESP', 'defender', 26, 9500000, 240000, 77, 79, 80, 0, false, 85, 86, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Luis', 'Martínez', 'ESP', 'defender', 27, 9800000, 242000, 79, 81, 78, 0, false, 86, 87, 88),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Manuel', 'Sánchez', 'ESP', 'defender', 28, 10200000, 248000, 81, 83, 80, 0, false, 88, 89, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Luca', 'Bianchi', 'ITA', 'defender', 29, 11500000, 270000, 83, 85, 82, 0, false, 89, 90, 91),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Mikkel', 'Jensen', 'DNK', 'midfielder', 27, 11000000, 265000, 82, 84, 81, 0, false, 88, 87, 89),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Antoine', 'Dubois', 'FRA', 'midfielder', 28, 11200000, 268000, 83, 85, 82, 0, false, 89, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Carlos', 'Silva', 'BRA', 'midfielder', 26, 11800000, 275000, 84, 86, 83, 0, false, 90, 89, 91),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Johan', 'Peterson', 'SWE', 'forward', 27, 11300000, 270000, 83, 84, 82, 0, false, 89, 88, 90),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), 'Pedro', 'Gonzalez', 'ARG', 'forward', 25, 11000000, 265000, 82, 83, 81, 0, false, 88, 87, 89);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Iván', 'Muñoz', 'ESP', 'goalkeeper', 30, 3200000, 90000, 72, 75, 70, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Mario', 'Ramos', 'ESP', 'defender', 29, 3000000, 88000, 70, 73, 75, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Javier', 'Castro', 'ESP', 'defender', 27, 2900000, 86000, 69, 71, 74, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Francisco', 'Ortiz', 'ESP', 'defender', 29, 3050000, 88000, 70, 73, 75, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Lars', 'Johansson', 'SWE', 'defender', 30, 3400000, 95000, 75, 77, 78, 0, false, 83, 85, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Mateo', 'Silva', 'BRA', 'midfielder', 28, 3200000, 90000, 72, 74, 70, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Raúl', 'Serrano', 'ESP', 'midfielder', 28, 3100000, 89000, 71, 74, 69, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Pablo', 'Marín', 'ESP', 'midfielder', 26, 2800000, 85000, 68, 70, 72, 0, false, 77, 79, 82),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Sergio', 'Torres', 'ESP', 'midfielder', 27, 3200000, 90000, 73, 75, 71, 0, false, 81, 83, 86),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'Diego', 'Vega', 'ESP', 'forward', 28, 3300000, 92000, 74, 76, 72, 0, false, 82, 84, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), 'John', 'Miller', 'USA', 'forward', 27, 3100000, 89000, 71, 73, 69, 0, false, 79, 81, 84);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Óscar', 'López', 'ESP', 'goalkeeper', 29, 3500000, 92000, 73, 76, 71, 0, false, 81, 83, 86),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Raúl', 'Molina', 'ESP', 'defender', 27, 3300000, 90000, 71, 74, 72, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Alberto', 'Herrera', 'ESP', 'defender', 28, 3400000, 91000, 72, 75, 73, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'David', 'Wilson', 'ENG', 'defender', 30, 3600000, 94000, 74, 77, 73, 0, false, 82, 84, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Felipe', 'Rodriguez', 'ARG', 'midfielder', 27, 3300000, 90000, 71, 74, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Iván', 'Ramos', 'ESP', 'midfielder', 26, 3100000, 88000, 70, 73, 69, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Miguel', 'Santos', 'ESP', 'midfielder', 27, 3200000, 89000, 71, 74, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Luis', 'Pérez', 'ESP', 'midfielder', 28, 3300000, 90000, 72, 75, 71, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'José', 'García', 'ESP', 'forward', 29, 3400000, 91000, 73, 76, 72, 0, false, 81, 83, 86),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Mark', 'Anderson', 'USA', 'forward', 28, 3400000, 91000, 72, 75, 71, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), 'Tom', 'Baker', 'AUS', 'forward', 29, 3500000, 92000, 73, 76, 72, 0, false, 81, 83, 86);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Carlos', 'Ruiz', 'ESP', 'goalkeeper', 30, 3400000, 91000, 73, 75, 72, 0, false, 81, 83, 86),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Javier', 'Fernández', 'ESP', 'defender', 29, 3300000, 90000, 72, 74, 71, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Manuel', 'Moreno', 'ESP', 'defender', 28, 3200000, 89000, 71, 73, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Francisco', 'Martínez', 'ESP', 'defender', 27, 3100000, 88000, 70, 72, 69, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Pedro', 'Ruiz', 'ESP', 'midfielder', 26, 3000000, 87000, 69, 71, 68, 0, false, 77, 79, 82),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Liam', 'Evans', 'WAL', 'midfielder', 28, 3200000, 89000, 71, 73, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Álvaro', 'Torres', 'ESP', 'midfielder', 27, 3100000, 88000, 70, 72, 69, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'David', 'Sánchez', 'ESP', 'midfielder', 26, 3000000, 87000, 69, 71, 68, 0, false, 77, 79, 82),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Miguel', 'López', 'ESP', 'forward', 28, 3200000, 89000, 71, 73, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Juan', 'Gómez', 'ESP', 'forward', 29, 3300000, 90000, 72, 74, 71, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), 'Andrés', 'Martinez', 'ARG', 'forward', 27, 3100000, 88000, 70, 72, 69, 0, false, 78, 80, 83);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary,
    technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Raúl', 'Cano', 'ESP', 'goalkeeper', 29, 3300000, 90000, 72, 75, 70, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Sergio', 'Morales', 'ESP', 'defender', 27, 3100000, 88000, 70, 72, 69, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Javier', 'Reyes', 'ESP', 'defender', 28, 3200000, 89000, 71, 73, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Erik', 'Larsen', 'DEN', 'defender', 30, 3500000, 94000, 74, 76, 73, 0, false, 82, 84, 87),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Luis', 'Santos', 'ESP', 'defender', 27, 3100000, 88000, 70, 72, 69, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Lucas', 'Martinez', 'ARG', 'midfielder', 28, 3300000, 90000, 71, 73, 70, 0, false, 79, 81, 84),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Miguel', 'Vargas', 'ESP', 'midfielder', 26, 3000000, 87000, 69, 71, 68, 0, false, 77, 79, 82),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Nicolás', 'Gómez', 'URU', 'midfielder', 28, 3300000, 90000, 72, 74, 71, 0, false, 80, 82, 85),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Marco', 'Rossi', 'ITA', 'midfielder', 27, 3200000, 89000, 70, 72, 69, 0, false, 78, 80, 83),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'James', 'Wilson', 'ENG', 'forward', 29, 3400000, 92000, 73, 75, 72, 0, false, 81, 83, 86),
(gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Río Claro'), 'Juan', 'Herrera', 'ESP', 'forward', 28, 3200000, 89000, 71, 73, 70, 0, false, 79, 81, 84);






INSERT INTO oft.strategy (id, team_id, formation, playing_style, game_tempo, passing_style, defensive_positioning, build_up_play, attack_focus, key_player_usage)
VALUES
    
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Deportivo Bahía Real'), '4-3-3', 'possession', 'fast_tempo', 'short', 'zonal_marking', 'play_from_back', 'wide_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Atlético Sierra Norte'), '3-4-3', 'high_press', 'balanced_tempo', 'short', 'man_marking', 'play_from_back', 'central_play', 'free_role_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Unión Deportiva Costa Verde'), '4-3-3', 'counter_attack', 'fast_tempo', 'long', 'zonal_marking', 'long_clearance', 'central_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Fútbol Club Valle Azul'), '5-3-2', 'direct_play', 'slow_tempo', 'long', 'man_marking', 'long_clearance', 'wide_play', 'free_role_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Club Atlético Rocafuerte'), '4-4-2', 'low_block', 'slow_tempo', 'short', 'zonal_marking', 'play_from_back', 'central_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Deportivo Villa del Mar'), '3-4-3', 'possession', 'balanced_tempo', 'short', 'zonal_marking', 'play_from_back', 'wide_play', 'free_role_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Agrupación Deportiva Puente Largo'), '3-4-3', 'high_press', 'fast_tempo', 'short', 'man_marking', 'play_from_back', 'wide_play', 'reference_player'),

    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Sporting Monteluz CF'), '4-4-2', 'direct_play', 'balanced_tempo', 'long', 'man_marking', 'long_clearance', 'central_play', 'free_role_player');





