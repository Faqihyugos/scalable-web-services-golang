package main

import (
	"assign-2/models"
	"assign-2/routers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title          API Assignment 2 Hacktiv8
// @version        1.0
// @description    Assignment 2 Go Programming Language
// @termsOfService http://swagger.io/terms/

// @contact.name  Order API CRUD
// @contact.url   http://www.swagger.io/support
// @contact.email faqihyugos@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /

var (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "orders_by"
	passwd = "postgres"
	db     *gorm.DB
	err    error
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable", host, port, user, passwd, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
	routers.StartServer(db).Run()
}
