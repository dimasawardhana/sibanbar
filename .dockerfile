FROM golang:latest AS go_builder
ADD .
WORKDIR .
run go mod download
run go get -u github.com/jinzhu/gorm
run go get -u github.com/joho/dotenv
run go get -u github.com/go-sql-driver/mysql
run go get -u github.com/gin-gonic/gin