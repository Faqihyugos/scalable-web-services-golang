# Assignment 2

## Useful Library

```
// Gin gonic (Framework for Rest API)
go get -u github.com/gin-gonic/gin

// Postgres Driver
go get -u github.com/lib/pq
go get -u "gorm.io/driver/postgres"

// Gorm - ORM Library
go get -u gorm.io/gorm

go get -u "github.com/jinzhu/gorm/dialects/mysql"

// Swaggo - Documentation
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template


```

## Basic Go Command

```
// Lihat versi
go version

// Inisialisasi go modules
go mod init <nama-project>

// Mengeksekusi file go
go run main.go


// Kompilasi file menjadi exe
go build -o program.exe

// Memvalidasi dependensi yang belum terdownload
go mod tidy

```