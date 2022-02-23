package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "com.ocbc.smb/database"
	"com.ocbc.smb/router"
	"com.ocbc.smb/utils"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {

	utils.Init()
	dbMigrate()
	serverPort := beego.AppConfig.DefaultString("httpport", "8888")
	runMode := beego.AppConfig.DefaultString("gin.mode", "debug")
	gin.SetMode(runMode)

	routersInit := router.InitRouter()
	routersInit.Run(":" + serverPort)
}

func dbMigrate() {

	dbuser := beego.AppConfig.DefaultString("db.postgres.user", "user_adm1n")
	dbpass := beego.AppConfig.DefaultString("db.postgres.pass", "_pa$sw0rd321.")
	dbaddres := beego.AppConfig.DefaultString("db.postgres.addres", "localhost")
	dbport := beego.AppConfig.DefaultString("db.postgres.port", "5432")
	dbname := "db_smb"

	filename := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", dbport, dbaddres, dbuser, dbpass, dbname)

	db, err := sql.Open("postgres", filename)
	if err != nil {
		utils.LogInfo("Cannot open database  1 !" + err.Error())
	}
	err2 := db.Ping()
	if err2 != nil {
		utils.LogInfo("Cannot open database  2 !" + err2.Error())
	}

	fmt.Println("You are Successfully connected!")

	// migrations := &migrate.PackrMigrationSource{
	// 	Box: packr.New("migrations", "./migration"),
	// }
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	utils.LogInfo("Applied " + strconv.Itoa(n) + " migrations!")

}
