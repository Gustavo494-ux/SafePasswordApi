
use SafePassword;

CREATE TABLE IF NOT exists Users(
	Id int primary key auto_increment,
    Name text, 
    Email text,
    Email_Hash varchar(128),
    SafePassword varchar(128), 
    created_at timestamp default current_timestamp
)ENGINE=InnoDB;


CREATE TABLE IF NOT exists Credenciais(
	Id int primary key auto_increment,
    UsuarioId int, 
    Descricao text,
    siteUrl text,
    Login text,
    Senha text,
    created_at timestamp default current_timestamp
)ENGINE=InnoDB;

alter table Credenciais add foreign key(UsuarioId) references Users(id) ON DELETE CASCADE;

CREATE INDEX idx_email_hash ON Users (Email_Hash);


