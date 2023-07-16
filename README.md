# SafePasswordApi

O SafePasswordApi é um gerenciador de senhas simples, projetado para armazenar e gerenciar suas credenciais de forma segura. O projeto é desenvolvido em Golang e utiliza criptografia AES e RSA para garantir a proteção dos dados confidenciais armazenados. 

## Executanto o projeto

Antes de iniciar o SafePasswordApi, certifique-se de ter o Go (Golang) e o Mysql instalados em seu ambiente de desenvolvimento. Caso ainda não tenha, faça a instalação adequada antes de prosseguir.
### Configuração manual
1. Clone o repositório do SafePasswordApi em seu ambiente local:
```makefile
git clone https://github.com/Gustavo494-ux/SafePasswordApi.git
```

2. Agora você precisa configurar as informações de conexão do banco de dados MySQL no arquivo `.env`, localizado no diretorio src. Abra o arquivo `.env` em um editor de texto e preencha os campos apropriados com as informações corretas. as informações abaixo são apenas exemplos:
```makefile
DB_USUARIO=root
DB_SENHA=Senha123
DB_NAME=SafePassword
HOST_DATABASE=localhost
Port_DATABASE=3306

API_PORT=8000
SecretKey=8fb29448faee18b656030e8f5a8
```
3. Abra seu SGBD e execute o script que pode ser encontrado em src/Init.SQL
4. Realize a importação do arquivo de configuração das rotas na sua ferramenta de teste de API. Utilizamos o insomnia para realizar o consumo da API.
5. Execute o projeto SafePasswordApi com o seguinte comando:
```makefile
   go run main.go
```

3. Abra seu gerenciador de banco de dados e execute o script que pode ser encontrado em sc/SQL/Criar.SQL
4. Realize a importação do arquivo de configuração das rotas na sua ferramenta de teste de API. Utilizamos o insomnia para realizar o consumo da API.
5. Abra o terminal ou o seu equivalente no sistema operacional e no diretorio raiz execute o comando:
```makefile
   go run main.go
```
Após a execução bem-sucedida, o SafePasswordApi estará rodando na porta definida em API_PORT no arquivo .env.
Agora você pode interagir com a API para gerenciar suas senhas de forma segura.

Observação: Lembre-se de manter as informações de conexão com o banco de dados seguras e não compartilhe com terceiros.
## Segurança
### Criptografia e confidencidencialidade dos dados
A segurança dos dados no SafePasswordApi é reforçada por uma abordagem de criptografia de duas camadas, que envolve o uso combinado de AES 256 e RSA. Essa combinação poderosa fornece uma proteção abrangente para as informações sensíveis armazenadas no gerenciador de senhas.

Quando um usuário insere suas credenciais ou qualquer outra informação confidencial, o SafePasswordApi aplica inicialmente a criptografia AES 256. Essa primeira camada de segurança utiliza o algoritmo AES 256 bits para criptografar os dados, o que torna a informação ilegível e ininteligível para qualquer pessoa que tente acessá-la sem a chave de descriptografia adequada. O AES 256 é altamente confiável e amplamente considerado como uma das formas mais seguras de criptografia disponíveis atualmente.

Além disso, o SafePasswordApi utiliza a criptografia RSA como segunda camada de proteção. O algoritmo RSA é usado para criptografar a chave AES 256 que foi usada para criptografar os dados originais. Essa abordagem de duas camadas, onde a chave AES 256 é criptografada pela chave RSA, adiciona uma camada adicional de segurança e autenticação. Essa técnica é extremamente eficaz para proteger as chaves de criptografia e garantir que apenas os usuários autorizados possam acessar e decifrar as informações confidenciais.

O uso combinado de AES 256 e RSA no SafePasswordApi é uma estratégia de segurança altamente sofisticada, projetada para proteger efetivamente os dados sensíveis dos usuários. Com essa abordagem de duas camadas, os usuários podem ter a tranquilidade de que suas informações confidenciais estão sendo armazenadas e gerenciadas com um alto nível de segurança, mesmo em cenários de acesso não autorizado. A combinação desses poderosos algoritmos de criptografia reforça a confiança dos usuários no SafePasswordApi como um gerenciador de senhas verdadeiramente seguro e confiável.

### Testes Unitários para Segurança e Confiabilidade
O SafePasswordApi é desenvolvido com uma abordagem centrada na segurança e qualidade do código. Para garantir o melhor funcionamento do projeto, são empregados testes unitários de forma abrangente. Os testes unitários validam minuciosamente o comportamento individual de cada componente do código, assegurando que suas funcionalidades sejam implementadas corretamente.

Ao utilizar testes unitários, verificamos detalhadamente a funcionalidade de cada parte do código, incluindo a robusta criptografia AES 256 e RSA, que protege os dados confidenciais dos usuários. Essa abordagem de testes isolados ajuda a manter o código organizado e modular, facilitando a manutenção e identificação de problemas rapidamente, caso surjam.

Com a utilização exclusiva de testes unitários, temos a certeza de que o SafePasswordApi funciona de maneira confiável, atendendo às expectativas dos usuários em relação à segurança e eficácia do gerenciador de senhas. Os testes unitários são uma parte fundamental de nosso compromisso em fornecer um produto seguro, confiável e de alta qualidade.
