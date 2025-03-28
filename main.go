package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
	"test/endpoint"
	"test/server_config"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	serverAttribute := server_config.NewServerAttribute()
	err := serverAttribute.Init()
	if err != nil {
		panic(err)
	}

	dbMigrate(serverAttribute.DBConnection)

	r := gin.Default()
	//r.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	//})

	endpoints := endpoint.NewEndpoint(
		r,
		serverAttribute.DBConnection,
		serverAttribute.Services.BookService,
		serverAttribute.Services.SessionService,
	)

	endpoint.InitEndpoints(&endpoints)

	r.Run(":8080")
}

func dbMigrate(db *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}
	if db != nil {
		_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			log.Fatalf("Failed to migrate DB: %v", err)
		} else {
			fmt.Println("success migrate db ")
		}
	} else {
		os.Exit(3)
	}
}
