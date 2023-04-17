# Golang-Crud
<div align="center">
<img src="https://rollingstone.uol.com.br/media/uploads/chorao-charlie-brown-jr_reproducao_instagram.jpg" width="500rem">
</div>
Crud feito em Golang, com o objetivo de estudos, fiz com a temática de Charlie Brown Jr, uma das minhas bandas favoritas de todos os tempos, e uma singela homenagem ao Chorão.

## Funcionalidades
A aplicação implementa os principais conceitos de um crud, os métodos de ler, inserir, atualizar e deletar dados, no caso, músicas do Charlie Brown Jr, com o tempo pretendo deixar o projeto um pouco mais robusto.

## Tecnologias Usadas
- Golang versão 1.20
- Gin
- Air
- Gorm
- Viper
- Docker

## Inicializando o Projeto
Primeiramente crie uma pasta chamda **envs** com arquivo .env dentro do diretório **pkg/common** e coloque as variáveis de ambiente de acordo com o seguinte modelo:
```
PORT=:8080 //porta onde a aplicação irá rodar na máquina
DB_URL=postgres://USUÁRIO:SENHA=@HOST:PORTA // string de conexão com o banco de dados, deve estar de acordo com o que foi definido no arquivo docker-compose.yml
```
Para inicializar o projeto basta ter o docker instalado em sua máquina, mas primeiramente, certifique-se que que não tem nenhum serviço do postgres em execução na porta **5234** e apó isso execute os comandos abaixo para subir o container da aplicação:

### Para criar uma imagem a partir do Dockerfile:
```
docker build .
```

### Servidor para desenvolvimento
Para subir o container em versão de desenvolvimento, execute o seguinte comando:
```
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```
Para derrubar o servidor:
```
docker compose -f docker-compose.yml -f docker-compose.dev.yml down -v
```

### Servidor para produção:
Para subir o container em versão de produção, execute o seguinte comando:
```
docker compose -f docker-compose.yml -f docker-compose.build.yml up -d
```
Para derrubar o servidor:
```
docker compose -f docker-compose.yml -f docker-compose.build.yml down
```
