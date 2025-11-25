CREATE TABLE family_list (
                             fl_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                             cst_id INT NOT NULL,
                             fl_relation VARCHAR(50) NOT NULL,
                             fl_name VARCHAR(50) NOT NULL,
                             fl_dob VARCHAR(50) NOT NULL,

                             CONSTRAINT fk_family_customer
                                 FOREIGN KEY (cst_id) REFERENCES customer(cst_id)
                                     ON UPDATE CASCADE
                                     ON DELETE CASCADE
);
