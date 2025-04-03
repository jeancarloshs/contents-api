contents-api/
│── cmd/               # Pontos de entrada do aplicativo
│   └── main.go        # Arquivo principal do projeto
│── internal/          # Código interno do projeto (não acessível externamente)
│   ├── handlers/      # Manipuladores (controllers)
│   ├── services/      # Regras de negócio (use cases)
│   ├── repository/    # Camada de acesso a dados (banco de dados)
│   ├── models/        # Estruturas e entidades do projeto
│   └── config/        # Configurações do projeto (ex: leitura de .env)
│── pkg/               # Pacotes reutilizáveis (opcional)
│── api/               # Definições da API (ex: OpenAPI/Swagger)
│── migrations/        # Scripts de migração do banco de dados
│── scripts/           # Scripts auxiliares (ex: inicialização do banco)
│── go.mod             # Arquivo de módulo do Go
│── go.sum             # Gerenciamento de dependências
│── Dockerfile         # Arquivo para Docker
│── .env               # Variáveis de ambiente
│── README.md          # Documentação do projeto


cmd/ → Contém os arquivos de inicialização da aplicação.

internal/ → Implementação principal do sistema, não deve ser acessível externamente.

handlers/ → Controllers responsáveis por processar as requisições HTTP.

services/ → Contém as regras de negócio da aplicação (ex: processamento de dados).

repository/ → Responsável pela comunicação com o banco de dados.

models/ → Define as estruturas de dados do projeto.

config/ → Configurações da aplicação, como conexão com o banco.

pkg/ → Pacotes reutilizáveis que podem ser usados por outros projetos.

api/ → Arquivos de definição de API, como Swagger/OpenAPI.

migrations/ → Scripts para criação e atualização do banco de dados.

scripts/ → Scripts auxiliares, como comandos para popular o banco.

https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md