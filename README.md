# golang-crud
<div align="center">
<img src="https://rollingstone.uol.com.br/media/uploads/chorao-charlie-brown-jr_reproducao_instagram.jpg" width="500rem">
</div>
Crud feito em golang, com o objetivo de estudos, fiz com a temática de Charlie Brown Jr, uma das minhas bandas favoritas de todos os tempos, uma singela homenagem ao Chorão.

## Funcionalidades
A aplicação implementa os principais conceitos de um crud, os métodos de ler, inserir, atualizar e deletar dados, no caso, músicas do Charlie Brown Jr, com o tempo pretendo deixar o projeto um pouco mais robusto.

## Tecnologias Usadas
- Golang versão 1.20
- Gin
- Gorm
- Viper
- Docker
## Inicializando o Projeto
Para inicializar o projeto basta ter a linguagem Golang versão 1.20 instalada em sua máquina e o docker
Depois de fazer o git clone execute o comando abaixo para baixar as dependências do projeto:
```
go mod tidy
```
Após isso, certifique-se que que não tem nenhum serviço do postgres em execução na porta 5234 e execute o comando abaixo para subir o container do banco de dados:
```
docker compose up -d
```
Depois crie uma pasta chamda "envs" com arquivo .env dentro do diretório "pkg/common" e coleque as variáveis de ambiente de acordo com o seguinte modelo:
```
PORT=:8080 //porta onde a aplicação irá rodar na máquina
DB_URL=postgres://USUÁRIO:SENHA=@HOST:PORTA // string de conexão com o banco de dados, deve estar de acordo com o que foi definido no arquivo docker-compose.yml
```
Após isso basta iniciar o servidor com o comando:
```
go run cmd/main.go
```
