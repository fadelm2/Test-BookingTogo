CREATE TABLE nationality (
                             nationality_id  INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                             nationality_name VARCHAR(50) NOT NULL,
                             nationality_code CHAR(2) NOT NULL
);
