CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS oft.player (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id UUID REFERENCES oft.team(id) ON DELETE CASCADE,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    nationality CHAR(3) NOT NULL REFERENCES oft.country(code),
    position VARCHAR(255) NOT NULL,
    age INT,
    fee INT,
    salary INT,
    technique INT,
    mental INT,
    physique INT,
    injurydays INT DEFAULT 0,
    lined BOOLEAN DEFAULT false,
    familiarity INT,
    fitness INT,
    happiness INT
);

INSERT INTO oft.player (id, team_id, firstname, lastname, nationality, position, age, fee, salary, technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Marc-André', 'ter Stegen', 'DEU', 'goalkeeper', 32, 35000000, 700000, 85, 88, 80, 0, false, 95, 90, 93),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Ronald', 'Araújo', 'URY', 'defender', 25, 60000000, 850000, 80, 85, 90, 0, false, 90, 92, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Jules', 'Koundé', 'FRA', 'defender', 26, 55000000, 800000, 82, 84, 85, 0, false, 88, 91, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Alejandro', 'Balde', 'ESP', 'defender', 21, 50000000, 750000, 83, 80, 85, 0, false, 85, 87, 90),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Frenkie', 'de Jong', 'NLD', 'midfielder', 27, 90000000, 1000000, 91, 89, 82, 0, false, 92, 90, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Ilkay', 'Gündogan', 'DEU', 'midfielder', 34, 30000000, 950000, 90, 88, 80, 0, false, 95, 87, 92),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Raphinha', 'Dias', 'BRA', 'forward', 28, 60000000, 900000, 88, 85, 86, 0, false, 90, 88, 93),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Robert', 'Lewandowski', 'POL', 'forward', 36, 45000000, 1100000, 92, 90, 85, 0, false, 97, 91, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'João', 'Felix', 'PRT', 'midfielder', 25, 70000000, 850000, 90, 84, 82, 0, false, 89, 85, 91),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Lionel', 'Messi', 'ARG', 'forward', 36, 50000000, 1000000, 95, 90, 80, 0, false, 100, 90, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'FC Barcelona'), 'Pedri', 'Gonzalez', 'ESP', 'midfielder', 21, 70000000, 800000, 90, 85, 75, 0, false, 90, 88, 92);

INSERT INTO oft.player (
    id, team_id, firstname, lastname, nationality, position, age, fee, salary, technique, mental, physique, injurydays, lined, familiarity, fitness, happiness)
VALUES
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Ederson', 'Moraes', 'BRA', 'goalkeeper', 31, 40000000, 800000, 84, 87, 85, 0, false, 95, 91, 93),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Kyle', 'Walker', 'GBR', 'defender', 34, 50000000, 850000, 80, 85, 92, 0, false, 90, 92, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Rúben', 'Dias', 'PRT', 'defender', 27, 70000000, 950000, 83, 88, 90, 0, false, 92, 91, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'John', 'Stones', 'GBR', 'defender', 30, 60000000, 900000, 82, 87, 85, 0, false, 89, 90, 95),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Joško', 'Gvardiol', 'HRV', 'defender', 24, 80000000, 1000000, 85, 86, 88, 0, false, 91, 90, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Rodri', 'Hernandez', 'ESP', 'midfielder', 29, 100000000, 1200000, 91, 90, 87, 0, false, 97, 93, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Bernardo', 'Silva', 'PRT', 'midfielder', 30, 90000000, 1100000, 93, 89, 82, 0, false, 95, 92, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Jack', 'Grealish', 'GBR', 'forward', 29, 70000000, 900000, 89, 85, 84, 0, false, 91, 89, 94),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Phil', 'Foden', 'GBR', 'forward', 25, 95000000, 1100000, 92, 87, 85, 0, false, 96, 93, 97),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Erling', 'Haaland', 'NOR', 'forward', 24, 150000000, 1200000, 88, 85, 95, 0, false, 95, 92, 96),
    (gen_random_uuid(), (SELECT id FROM oft.team WHERE name = 'Manchester City'), 'Kevin', 'De Bruyne', 'BEL', 'midfielder', 33, 80000000, 1100000, 93, 90, 85, 0, false, 100, 90, 94);
