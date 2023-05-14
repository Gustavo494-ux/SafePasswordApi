
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
    Nome varchar(30),
    Senha varchar(100),
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;

CREATE TABLE IF NOT exists Senhas_Usuario(
	Id int primary key auto_increment,
    Id_Usuario int, 
    Id_Senhas int
)ENGINE=InnoDB;
alter table Senhas_Usuario add foreign key(Id_Usuario) references Usuarios(id);
alter table Senhas_Usuario add foreign key(Id_Senhas) references Senhas(id);


CREATE TABLE IF NOT exists Usuarios_Ativos(
	Id int primary key auto_increment,
    Id_Usuario int, 
    Id_Senhas int,
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;
alter table Usuarios_Ativos add foreign key(Id_Usuario) references Usuarios(id);