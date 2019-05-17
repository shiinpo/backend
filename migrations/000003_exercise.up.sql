CREATE TABLE IF NOT EXISTS exercise (
  id                serial          PRIMARY KEY,
  name              varchar(80)     UNIQUE NOT NULL,
  category_id       INTEGER         REFERENCES category(id) ON DELETE CASCADE ON UPDATE CASCADE
);