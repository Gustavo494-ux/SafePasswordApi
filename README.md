# SafePasswordApi

O SafePasswordApi é um gerenciador de senhas simples, projetado para armazenar e gerenciar suas credenciais de forma segura. O projeto é desenvolvido em Golang e utiliza criptografia AES (Advanced Encryption Standard) para garantir a proteção dos dados confidenciais armazenados. 

## Executanto o projeto

Antes de iniciar o SafePasswordApi, certifique-se de ter o Go (Golang) instalado em seu ambiente de desenvolvimento. Caso ainda não tenha, faça a instalação adequada antes de prosseguir.

Após ter o Go instalado, siga as etapas abaixo para configurar e executar o projeto:

1. Clone o repositório do SafePasswordApi em seu ambiente local:
```makefile
git clone https://github.com/Gustavo494-ux/SafePasswordApi.git
```

2. Agora você precisa configurar as informações de conexão do banco de dados MySQL no arquivo `.env`, localizado no diretorio src. Abra o arquivo `.env` em um editor de texto e preencha os campos apropriados com as informações corretas:
```makefile
DB_USUARIO=root
DB_SENHA=Senha123
DB_NAME=SafePassword
HOST_DATABASE=localhost
Port_DATABASE=3306

API_PORT=8000
SecretKey=8fb29448faee18b656030e8f5a8
```

3. Abra seu SGBD e execute o script que pode ser encontrado em sc/SQL/Criar.SQL
4. Realize a importação do arquivo de configuração das rotas na sua ferramenta de teste de API. Utilizamos o insomnia para realizar o consumo da API.
5. Execute o projeto SafePasswordApi com o seguinte comando:
```makefile
   go run main.go
```
Após a execução bem-sucedida, o SafePasswordApi estará rodando localmente na porta definida em API_PORT no arquivo .env.
Agora você pode interagir com a API para gerenciar suas senhas de forma segura.

Observação: Lembre-se de manter as informações de conexão com o banco de dados seguras e não compartilhe com terceiros.
Espero que isso atenda às suas necessidades! Se você tiver mais perguntas, estou aqui para ajudar.