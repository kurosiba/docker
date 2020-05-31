CREATE DATABASE gameapi_db;
use gameapi_db;

CREATE TABLE user (
    id integer(10), 
    username varchar(255), 
    token varchar(255)); 


INSERT INTO user (id, username, token) VALUES ( 1, 'kawauchiya' , 1);
