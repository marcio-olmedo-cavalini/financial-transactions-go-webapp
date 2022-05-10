Pacotes para instalar

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get gopkg.in/validator.v2
go get golang.org/x/crypto/bcrypt
=========================================================

Como executar:
1- Subir os containers do postgres: docker compose up
2- Executar o projeto: go run main.go

Assim que o projeto subir, executar no navegador: localhost:8080/index