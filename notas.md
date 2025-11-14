Orientações - Desafio Técnico (Simulação)

VERITAS CONSULTORIA EMPRESARIAL

Desafio Fullstack - App de Notas Pessoais (React + Go)

O objetivo deste desafio é avaliar a sua capacidade de desenvolver uma aplicação fullstack segura, implementando autenticação de usuários e controle de acesso a dados privados utilizando React no frontend e Go no backend.

Você deverá implementar um sistema onde usuários podem se cadastrar, fazer login e gerenciar suas próprias notas em Markdown.

Escopo Mínimo (MVP)

Frontend (React):

    Autenticação: Implementar telas de Login e Cadastro (Register).

    Sessão: Gerenciar a persistência do token (ex: localStorage) e redirecionar usuários não autenticados para o login.

    Dashboard: Renderizar a lista de notas pertencentes apenas ao usuário logado.

    Editor: Interface dividida (Split View) com área de texto para Markdown e visualização (preview) renderizada em tempo real.

    Feedbacks: Mensagens visuais de sucesso e erro (ex: "Usuário criado", "Senha inválida").

Backend (Go):

    Autenticação:

        Endpoint POST /register: Receber e-mail/senha e salvar o usuário com a senha criptografada (hash).

        Endpoint POST /login: Validar credenciais e retornar um token de sessão (JWT).

    Middleware: Implementar proteção nas rotas de notas, validando o token antes de processar a requisição.

    Endpoints RESTful: GET, POST, PUT e DELETE para /notes.

    Regra de Negócio: Garantir que as operações de CRUD afetem apenas as notas vinculadas ao ID do usuário autenticado.

    CORS: Configurado para permitir acesso do frontend.

Estrutura de Entrega

O projeto deve ser organizado em um repositório no GitHub da seguinte forma:
Plaintext

/backend
main.go, handlers.go, models.go
auth.go, middleware.go (sugestão)
/frontend
package.json, src/...
/docs
user-flow.png (obrigatório)
auth-flow.png (opcional: fluxo de login)
README.md

O README deve conter:

    Instruções para rodar o backend e o frontend.

    Explicação sobre a estratégia de autenticação (bibliotecas usadas, expiração do token).

    Decisões técnicas tomadas.

Critérios de Avaliação

    Segurança (30%) - Senhas com hash, validação correta de JWT e proteção de rotas.

    Funcionalidade (30%) - Fluxo de login/logout e CRUD privado funcionando perfeitamente.

    Código e Arquitetura (20%) - Organização, separação de responsabilidades (middleware vs handlers).

    Documentação (15%) - Clareza no README e User Flow.

    UI/UX (5%) - Experiência de escrita fluida.

Bônus (Opcional)

    Infra: Docker e Docker Compose para todo o ambiente.

    UX: Funcionalidade de Auto-save (salvar rascunho automaticamente).

    Persistência Real: Uso de banco de dados SQL (SQLite ou Postgres) ao invés de memória/JSON.

    Testes: Testes unitários na lógica de autenticação.

Instruções Finais

Crie um repositório público no GitHub com o nome desafio-fullstack-notas, siga as instruções deste documento e trate-o como um projeto real de portfólio. A prioridade é a segurança e o funcionamento correto da segregação de dados por usuário.
