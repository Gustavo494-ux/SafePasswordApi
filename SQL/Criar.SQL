drop database SafePassword;

create database if not exists SafePassword;
use SafePassword;

CREATE TABLE IF NOT exists Usuarios(
	Id int primary key auto_increment,
    Nome varchar(100), 
    Email varchar(100),
    Senha varchar(100), 
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;

CREATE TABLE IF NOT exists Senhas(
	Id int primary key auto_increment,
    UsuarioId int, 
    Descricao varchar(30),
    Login varchar(100),
    Senha varchar(100),
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;

alter table Senhas add foreign key(UsuarioId) references Usuarios(id) ON DELETE CASCADE;
