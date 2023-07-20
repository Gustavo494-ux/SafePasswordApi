# SafePassword

O SafePassword é um gerenciador de senhas simples, projetado para armazenar e gerenciar suas credenciais de forma segura. O projeto é desenvolvido em Golang e utiliza criptografia AES e RSA para garantir a proteção dos dados confidenciais armazenados.

## Bancos suportados
- MySQL

## Executanto o projeto

Antes de iniciar o SafePassword, certifique-se de ter o Go (Golang) e o MySQL instalados em seu ambiente. Caso ainda não tenha, faça a instalação adequada antes de prosseguir.

1. Clone o repositório do SafePassword em seu ambiente local:
```makefile
git clone https://github.com/Gustavo494-ux/SafePasswordApi
```


2. Configure as informações de conexão do banco de dados MySQL no arquivo `.env`, localizado no diretório `SafePassworApi/src`. Preencha os campos apropriados com as informações corretas. As informações abaixo são apenas exemplos:
```makefile
DB_USER=root
DB_PASSWORD=Senha123
DB_NAME=SafePassword
DB_HOST=localhost
DB_PORT=3306

API_PORT=8000

RSA_Private_key_Path="C:/Projetos/Keys/RSA_Private_Key.txt"
RSA_Public_key_Path="C:/Projetos/Keys/RSA_Public_Key.txt"
AES_key_Path="C:/Projetos/Keys/AES_Key.txt"
SECRET_KEY_JWT_PATH = "C:/Projetos/Keys/SECRET_KEY_JWT_PATH.txt"
```
As chaves utilizadas para encriptação ou autenticação são geradas automaticamente ao executar ou realizar testes no projeto. Recomenda-se que deixe o SafePassword gerar todas as chaves necessárias e depois seja realizado backup delas em um lugar seguro. 
pois caso a chave AES ou chave privada RSA seja perdida não será possível reverter a encriptação dos dados armazenados.

1. Abra seu gerenciador de banco de dados e execute o script que pode ser encontrado em sc/SQL/Criar.SQL.
2. Realize a importação do arquivo de configuração das rotas na sua ferramenta de teste de API. Utilizamos o Insomnia, mas pode utilizar o que mais lhe agradar.
3. Abra o terminal ou o seu equivalente no sistema operacional e no diretório raiz execute o comando:
```makefile
   go run main.go
```
Após a execução bem-sucedida, o SafePassword estará rodando na porta definida em API_PORT no arquivo .env. Agora você pode interagir com a API para gerenciar suas senhas de forma segura.

Observação: Lembre-se de manter as informações de conexão com o banco de dados seguras e não compartilhe com terceiros.

### Comandos úteis e suas funções
Todos os comandos abaixo deve ser executado no terminal ou seu equivalente no sistema operacional. e sempre no diretorio raiz(onde está localizado o main.go) do projeto

1. Executa o projeto
```
go run main.go
```
2. Realiza o Build do projeto, assim gerando um exeutável.
```
go build
```
3. Executa os testes.
```
go test
```
## Segurança
### Criptografia e confidencidencialidade dos dados
A segurança dos dados no SafePassword é reforçada por uma abordagem de criptografia de duas camadas, que envolve o uso combinado de AES 256 e RSA. Essa combinação poderosa fornece uma proteção abrangente para as informações sensíveis armazenadas no gerenciador de senhas.

Quando um usuário insere suas credenciais ou qualquer outra informação confidencial, o SafePassword aplica inicialmente a criptografia AES 256. Essa primeira camada de segurança utiliza o algoritmo AES 256 bits para criptografar os dados, o que torna a informação ilegível e ininteligível para qualquer pessoa que tente acessá-la sem a chave de descriptografia adequada. O AES 256 é altamente confiável e amplamente considerado como uma das formas mais seguras de criptografia disponíveis atualmente.

Além disso, o SafePassword utiliza a criptografia RSA como segunda camada de proteção. O algoritmo RSA é usado para criptografar a chave AES 256 que foi usada para criptografar os dados originais. Essa abordagem de duas camadas, onde a chave AES 256 é criptografada pela chave RSA, adiciona uma camada adicional de segurança e autenticação. Essa técnica é extremamente eficaz para proteger as chaves de criptografia e garantir que apenas os usuários autorizados possam acessar e decifrar as informações confidenciais.

O uso combinado de AES 256 e RSA no SafePassword é uma estratégia de segurança altamente sofisticada, projetada para proteger efetivamente os dados sensíveis dos usuários. Com essa abordagem de duas camadas, os usuários podem ter a tranquilidade de que suas informações confidenciais estão sendo armazenadas e gerenciadas com um alto nível de segurança, mesmo em cenários de acesso não autorizado. A combinação desses poderosos algoritmos de criptografia reforça a confiança dos usuários no SafePassword como um gerenciador de senhas verdadeiramente seguro e confiável.

### Testes Unitários para Segurança e Confiabilidade
O SafePassword é desenvolvido com uma abordagem centrada na segurança e qualidade do código. Para garantir o melhor funcionamento do projeto, são empregados testes unitários de forma abrangente. Os testes unitários validam minuciosamente o comportamento individual de cada componente do código, assegurando que suas funcionalidades sejam implementadas corretamente.

Ao utilizar testes unitários, verificamos detalhadamente a funcionalidade de cada parte do código, incluindo a robusta criptografia AES 256 e RSA, que protege os dados confidenciais dos usuários. Essa abordagem de testes isolados ajuda a manter o código organizado e modular, facilitando a manutenção e identificação de problemas rapidamente, caso surjam.

Com a utilização exclusiva de testes unitários, temos a certeza de que o SafePassword funciona de maneira confiável, atendendo às expectativas dos usuários em relação à segurança e eficácia do gerenciador de senhas. Os testes unitários são uma parte fundamental de nosso compromisso em fornecer um produto seguro, confiável e de alta qualidade.

## Documentação das rotas

### Login
#### Verbo HTTP POST. Não requer autenticação
##### A URL é composta pelo servidor:api_port do .env/Login por exemplo localhost:8000/Login. Abaixo segue um exemplo de JSON que deve ser enviado na requisição.
```
{
	"email": "usuario1@gmail.com",
	"password":"123456"
}
```
Caso esteja tudo correto será retornado um statusCode 202, e um token JWT que deverá ser informado em todas as requisições que a rota exiga autenticação. Caso esteja utilizando o Insomnia informe o token ao realizar a requisiçao no seguinte local em Auth/Bearer Token/Token, é fundamental que a opção Enabled esteja marcada.


### Usuário
#### Criar Usuário 
##### Verbo HTTP POST. Não requer autenticação
##### A URL é composta pelo servidor:api_port do .env/user por exemplo localhost:8000/user. Abaixo segue um exemplo de JSON que deve ser enviado na requisição
```
{
	"name":"Usuario 1",
	"email": "usuario1@gmail.com",
	"password":"123456"
}
```

#### Buscar Usuários
##### Verbo HTTP GET. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/users`` por exemplo ``localhost:8000/users``. Abaixo segue um não é necessario enviar o JSON.  Exemplo de resposta
```
[
	{
		"id": 10,
		"name": "Usuario 1",
		"email": "usuario1@gmail.com",
		"created_at": "2023-07-19T08:29:41Z"
	},
	{
		"id": 11,
		"name": "Usuario 1",
		"email": "usuario1@gmail.com",
		"created_at": "2023-07-20T21:36:59Z"
	},
	{
		"id": 12,
		"name": "Usuario 1",
		"email": "usuario1@gmail.com",
		"created_at": "2023-07-20T21:46:40Z"
	}
]
```

#### Buscar Usuário por Id
##### Verbo HTTP GET. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/users/usuarioId`` por exemplo ``localhost:8000/users/5``. Abaixo segue um não é necessario enviar o JSON.  Exemplo de resposta
```
{
	"id": 5,
	"name": "Usuario 1",
	"email": "usuario1@gmail.com",
	"created_at": "2023-07-16T15:29:47Z"
}
```

#### Atualizar Usuário
##### Verbo HTTP PUT. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/users/usuarioId`` por exemplo ``localhost:8000/users/5``. Abaixo   exemplo de JSON que deve ser enviado na requisição
```
{
	"name":"Usuario 3",
	"email": "usuario2@gmail.com",
	"password":"123456"
}
```
Caso esteja tudo certo esperasse um status code 200. 

#### Deletar Usuário
##### Verbo HTTP Delete. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/users/usuarioId`` por exemplo ``localhost:8000/users/5``. Abaixo   Caso esteja tudo certo esperasse um status code 204. 


### Credencial
#### Criar Credencial 
##### Verbo HTTP POST. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/credentials`` por exemplo ``localhost:8000/credentials``. Abaixo segue um exemplo de JSON que deve ser enviado na requisição:
```
{
	"descricao":"Facebook",
	"siteUrl":"https://www.facebook.com/",
	"login":"usuario1@gmail.com",
	"senha":"Senha123"
}
```
Retorno: Status code 201 e a credencial. por exemplo:
```
{
		"id": 3,
		"usuarioId": 11,
		"descricao": "Facebook",
		"siteUrl": "https://www.facebook.com/",
		"login": "usuario1@gmail.com",
		"senha": "Senha123",
		"criadoEm": "2023-07-20T22:17:20Z"
	}
```


#### Buscar Credenciais
##### Verbo HTTP GET. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/credentials`` por exemplo ``localhost:8000/credentials``. Abaixo segue um não é necessario enviar o JSON.  Exemplo de resposta
```
[
	{
		"id": 3,
		"usuarioId": 11,
		"descricao": "Facebook",
		"siteUrl": "https://www.facebook.com/",
		"login": "usuario1@gmail.com",
		"senha": "Senha123",
		"criadoEm": "2023-07-20T22:17:20Z"
	},
	{
		"id": 4,
		"usuarioId": 11,
		"descricao": "Facebook",
		"siteUrl": "https://www.facebook.com/",
		"login": "usuario1@gmail.com",
		"senha": "Senha123",
		"criadoEm": "2023-07-20T22:17:21Z"
	}
]
```

#### Buscar Credenciail por Id
##### Verbo HTTP GET. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/credentials/credencilId`` por exemplo ``localhost:8000/credentials/3``. Abaixo segue um não é necessario enviar o JSON.  Exemplo de resposta
```
{
   "id": 3,
   "usuarioId": 11,
   "descricao": "Facebook",
   "siteUrl": "https://www.facebook.com/",
   "login": "usuario1@gmail.com",
   "senha": "Senha123",
   "criadoEm": "2023-07-20T22:17:20Z"
}
```

#### Atualizar Credencial
##### Verbo HTTP PUT. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/credentials/credencialId`` por exemplo ``localhost:8000/credentials/5``. Abaixo exemplo de JSON que deve ser enviado na requisição
```
{
	"descricao":"Youtube",
	"siteUrl":"https://www.youtube.com/",
	"login":"usuarios1@gmail.com",
	"senha":"Senha123456"
}
```
Caso esteja tudo certo esperasse um status code 200. 

#### Deletar Credencial
##### Verbo HTTP Delete. Requer autenticação
##### A URL é composta pelo ``servidor:api_port do .env/credentials/credencialId`` por exemplo ``localhost:8000/credentials/5``. Abaixo   Caso esteja tudo certo esperasse um status code 204. 
