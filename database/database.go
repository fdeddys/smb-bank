package database

import (
	"fmt"

	"com.ocbc.smb/utils"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Dbcon    *gorm.DB
	Errdb    error
	dbuser   string
	dbpass   string
	dbname   string
	dbaddres string
	dbport   string
	dbdebug  bool
	dbtype   string
	sslmode  string
)

func init() {

	dbtype = beego.AppConfig.DefaultString("db.type", "POSTGRES")
	dbuser = beego.AppConfig.DefaultString("db.postgres.user", "user_adm1n")
	dbpass = beego.AppConfig.DefaultString("db.postgres.pass", "_pa$sw0rd321.")
	dbaddres = beego.AppConfig.DefaultString("db.postgres.addres", "localhost")
	dbport = beego.AppConfig.DefaultString("db.postgres.port", "5432")
	dbdebug = beego.AppConfig.DefaultBool("db.postgres.debug", true)
	dbname = "db_smb"
	sslmode = "disable"
	if DbOpen() != nil {
		fmt.Println("Cannot Open db Postgres !")
	}
}

// DbOpen ...
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbaddres, dbport, dbuser, dbpass, dbname, sslmode)
	Dbcon, Errdb = gorm.Open("postgres", args)
	// fmt.Println("isi postgres sett ", args)
	if Errdb != nil {
		utils.LogInfo("open db Err : " + Errdb.Error())
		return Errdb
	}
	if errping := Dbcon.DB().Ping(); errping != nil {
		return errping
	}
	fmt.Println("Database connected [", dbaddres, "] [", dbname, "] [", dbuser, "] !")
	return nil
}

// GetDbCon ...
func GetDbCon() *gorm.DB {

	if errping := Dbcon.DB().Ping(); errping != nil {
		utils.LogInfo("Db Not Connected, test Ping :" + errping.Error())
		errping = nil
		if errping = DbOpen(); errping != nil {
			utils.LogInfo("try to connect again, but error :" + errping.Error())
		}
	}
	Dbcon.LogMode(dbdebug)
	return Dbcon
}
