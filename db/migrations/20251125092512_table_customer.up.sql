CREATE TABLE customer (
                          cst_id  INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                          nationality_id INT NOT NULL,
                          cst_name CHAR(50) NOT NULL,
                          cst_dob DATE NOT NULL,
                          cst_phoneNum VARCHAR(20) NOT NULL,
                          cst_email VARCHAR(50) NOT NULL,

                          CONSTRAINT fk_customer_nationality
                              FOREIGN KEY (nationality_id) REFERENCES nationality(nationality_id)
                                  ON UPDATE CASCADE
                                  ON DELETE RESTRICT
);
