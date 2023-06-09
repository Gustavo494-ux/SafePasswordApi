drop database SafePassword;

create database if not exists SafePassword;
use SafePassword;

CREATE TABLE IF NOT exists Usuarios(
	Id int primary key auto_increment,
    Nome varchar(100), 
    Email varchar(100),
    Senha varchar(128), 
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;

INSERT INTO SafePassword.Usuarios (Nome,Email,Senha) values

('Usuario 1', 'usuario1@gmail.com', '$2a$10$s05Q2pymPXvM/KegOjraAOFYHlVJYUL3kSOyPJNXIxHon/2PKDY16'),
('Usuario 2', 'usuario2@gmail.com', '$2a$10$MGk.0Q9Op3DMcuJzoHAcce13r89UROdgNJ281CUkAFw5P2KRQks7W'),
('Usuario 3', 'usuario3@gmail.com', '$2a$10$lkraNpCrLoJ.tWoeFjdjaehLfEQeptyEWYDNJjgrqTX0BPPLnAhhK');