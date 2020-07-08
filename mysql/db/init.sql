CREATE DATABASE test_db;
use test_db;

CREATE TABLE users (
    username varchar(255),
    token integer(10)
    ); 

INSERT INTO users (username, token) VALUES ('example1', 1);
INSERT INTO users (username, token) VALUES ('example2', 2);
INSERT INTO users (username, token) VALUES ('example3', 3);