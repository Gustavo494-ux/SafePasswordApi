CREATE DATABASE IF NOT EXISTS devbook;
use devbook;

DROP DATABASE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null, 
    nick varchar(50) not null,
    email varchar(50) not null,
    senha varchar(20) not null,
    criadoem timestamp default current_timestamp()
)ENGINE=InnoDB;