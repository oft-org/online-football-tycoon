BEGIN;

CREATE TABLE oft.players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id UUID REFERENCES oft.teams(id) ON DELETE CASCADE,
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

COMMIT;
