
use SafePassword;

CREATE TABLE IF NOT exists Usuarios(
	Id int primary key auto_increment,
    Nome text, 
    Email text,
    Email_Hash varchar(128),
    Senha varchar(128), 
    criadoEm timestamp default current_timestamp
)ENGINE=InnoDB;


CREATE TABLE IF NOT exists Credenciais(
	Id int primary key auto_increment,
    UsuarioId int, 
    Descricao text,
    siteUrl text,
    Login text,
    Senha text,
    criadoEm timestamp default current_timestamp
)ENGINE=InnoDB;

alter table Credenciais add foreign key(UsuarioId) references Usuarios(id) ON DELETE CASCADE;

CREATE INDEX idx_email_hash ON Usuarios (Email_Hash);


