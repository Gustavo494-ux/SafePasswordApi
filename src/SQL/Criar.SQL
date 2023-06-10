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


CREATE TABLE IF NOT exists Credenciais(
	Id int primary key auto_increment,
    UsuarioId int, 
    Descricao varchar(30),
    siteUrl varchar(2048),
    Login varchar(100),
    Senha varchar(300),
    CriadoEm timestamp default current_timestamp
)ENGINE=InnoDB;

alter table Credenciais add foreign key(UsuarioId) references Usuarios(id) ON DELETE CASCADE;



INSERT INTO SafePassword.Usuarios (Nome,Email,Senha) values
    ('Usuario 1', 'usuario1@gmail.com', 'ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413'),
    ('Usuario 2', 'usuario2@gmail.com', 'ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413'),
    ('Usuario 3', 'usuario3@gmail.com', 'ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413');

INSERT INTO SafePassword.Credenciais ( UsuarioId,Descricao,Login,Senha) values
    (1,'Alura','https://www.alura.com.br/','usuario1@gmail.com','Senha1'),
    (1,'Udemy','https://www.udemy.com/','usuario_1','Senha2'),
    (1,'Youtube','usuario1_3833@empresa.com','Senha3'),
    (2,'Facebook','','','Senha1'),
    (2,'Spotify','','','Senha2');

