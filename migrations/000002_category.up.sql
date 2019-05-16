CREATE TABLE IF NOT EXISTS category (
  id        serial          PRIMARY KEY,
  name      varchar(40)     UNIQUE NOT NULL
);