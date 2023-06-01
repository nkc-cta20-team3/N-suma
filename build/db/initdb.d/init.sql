DROP SCHEMA IF EXISTS database;
CREATE SCHEMA database;
USE database;

DROP TABLE IF EXISTS table;

CREATE TABLE table
(
  id           INT(10),
  name     VARCHAR(40)
);

INSERT INTO table (id, name) VALUES (1, "Nagaoka");
INSERT INTO table (id, name) VALUES (2, "Tanaka");
