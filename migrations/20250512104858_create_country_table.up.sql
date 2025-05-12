BEGIN;
CREATE TABLE oft.country (
    code CHAR(3) PRIMARY KEY,
    continent TEXT NOT NULL,
    FOREIGN KEY (continent) REFERENCES oft.continent(code)
);

INSERT INTO oft.country (code, continent) VALUES 
  ('ARG', 'SOUTH_AMERICA'),  
  ('BRA', 'SOUTH_AMERICA'),  
  ('FRA', 'EUROPE'),         
  ('DEU', 'EUROPE'),         
  ('ESP', 'EUROPE'),         
  ('ITA', 'EUROPE'),         
  ('ENG', 'EUROPE'),         
  ('POR', 'EUROPE'),         
  ('NLD', 'EUROPE'),         
  ('BEL', 'EUROPE'),         
  ('URU', 'SOUTH_AMERICA'),  
  ('COL', 'SOUTH_AMERICA'),  
  ('CHL', 'SOUTH_AMERICA'),  
  ('MEX', 'NORTH_AMERICA'),  
  ('USA', 'NORTH_AMERICA'),  
  ('CIV', 'AFRICA'),         
  ('SEN', 'AFRICA'),         
  ('MAR', 'AFRICA'),         
  ('CMR', 'AFRICA'),         
  ('GHA', 'AFRICA'),         
  ('JPN', 'ASIA'),           
  ('KOR', 'ASIA'),           
  ('SAU', 'ASIA'),           
  ('IRN', 'ASIA'),           
  ('AUS', 'OCEANIA'),        
  ('SUI', 'EUROPE'),         
  ('CRO', 'EUROPE'),         
  ('POL', 'EUROPE'),         
  ('SRB', 'EUROPE'),         
  ('TUR', 'EUROPE'),         
  ('ECU', 'SOUTH_AMERICA'),  
  ('PER', 'SOUTH_AMERICA'),  
  ('ALG', 'AFRICA'),         
  ('QAT', 'ASIA'),           
  ('CAN', 'NORTH_AMERICA')
  ('NOR', 'EUROPE');  

COMMIT;
