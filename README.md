# SafePassword

O SafePassword é um gerenciador de senhas simples, projetado para armazenar e gerenciar suas credenciais de forma segura. O projeto é desenvolvido em Golang e utiliza criptografia AES e RSA para garantir a proteção dos dados confidenciais armazenados. 
#### Bancos suportados: Mysql
## Executanto o projeto

Antes de iniciar o SafePassword, certifique-se de ter o Go (Golang) e o Mysql instalados em seu ambiente. Caso ainda não tenha, faça a instalação adequada antes de prosseguir.
1. Clone o repositório do SafePassword em seu ambiente local:
```makefile
git clone https://github.com/Gustavo494-ux/SafePasswordApi
```

2. Agora você precisa configurar as informações de conexão do banco de dados MySQL no arquivo `.env`, localizado no diretorio  SafePassworApi/src. Abra o arquivo `.env` em um editor de texto e preencha os campos apropriados com as informações corretas. as informações abaixo são apenas exemplos:
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

3. Abra seu gerenciador de banco de dados e execute o script que pode ser encontrado em sc/SQL/Criar.SQL
4. Realize a importação do arquivo de configuração das rotas na sua ferramenta de teste de API. Utilizamos o insomnia, mas pode utilizar o que mais lhe agradar.
5. Abra o terminal ou o seu equivalente no sistema operacional e no diretorio raiz execute o comando:
```makefile
   go run main.go
```
Após a execução bem-sucedida, o SafePassword estará rodando na porta definida em API_PORT no arquivo .env.
Agora você pode interagir com a API para gerenciar suas senhas de forma segura.

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
