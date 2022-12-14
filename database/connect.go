package database

import (
	"TerraInnAPI/model"
	"fmt"
	"os"

	// postgresDriver "gorm.io/driver/postgres"

	"TerraInnAPI/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// var Store *postgres.Storage

func Connect() bool {
	var err error
	status := true
	//Var
	db_host := "db4free.net"
	db_port := "3306"
	db_user := "quyen_dvt"
	db_password := "VEKdEzf2e9EzdE"
	db_name := "terrainn_db"
	db_clearall := "true"
	db_init := "true"
	db_debugmode := "true"
	db_migrate := "true"
	//Check enviroment mode:
	if os.Getenv("PRODUCTION") == "true" {
		fmt.Println("**** ENV MODE = PRODUCTION ****")

		//.HARDCODE MODE:
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USER")
		db_password = os.Getenv("DB_PASSWORD")
		db_name = os.Getenv("DB_NAME")
		db_clearall = os.Getenv("DB_CLEARALL")
		db_init = os.Getenv("DB_INIT")
		db_debugmode = os.Getenv("DEBUG_MODE")
		db_migrate = os.Getenv("DB_MIGRATE")
	} else {
		fmt.Println("**** ENV MODE = DEVELOP ****")

		//.ENV MODE:
		db_host = config.Config("DB_HOST")
		db_port = config.Config("DB_PORT")
		db_user = config.Config("DB_USER")
		db_password = config.Config("DB_PASSWORD")
		db_name = config.Config("DB_NAME")
		db_clearall = config.Config("DB_CLEARALL")
		db_init = config.Config("DB_INIT")
		db_debugmode = config.Config("DEBUG_MODE")
		db_migrate = config.Config("DB_MIGRATE")

	}

	// db_ssh := config.Config("DB_SSH")
	// db_timezone := config.Config("APP_TIME_ZONE")

	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db_host, db_port, db_user, db_password, db_name, db_ssh)

	//MySQL DNS:
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)

	var configOptions gorm.Config
	// if config.Config("DEBUG_MODE") == "true" {
	if db_debugmode == "true" {

		configOptions = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //Debug ONLY
		}
	}
	DB, err = gorm.Open(mysql.Open(dsn), &configOptions)
	if err != nil {
		fmt.Println("CONNECT DB ERROR: ")
		panic(err)
	}
	// Open and create db: https://golangbot.com/connect-create-db-mysql/

	if !status {
		return false
	}

	//Session Config:
	// ConfigSession()

	/**
	Remove table
	Only use for dev
	*/

	if db_clearall == "true" {
		DB.Migrator().DropTable(&model.User{})
		DB.Migrator().DropTable(&model.Motel{})
		DB.Migrator().DropTable(&model.MotelGroup{})
		DB.Migrator().DropTable(&model.Invoice{})
		DB.Migrator().DropTable(&model.Request{})

	}

	if db_migrate == "true" {
		// Migrate the database
		DB.AutoMigrate(&model.User{})
		DB.AutoMigrate(&model.Motel{})
		DB.AutoMigrate(&model.MotelGroup{})
		DB.AutoMigrate(&model.Invoice{})
		DB.AutoMigrate(&model.Request{})
	}

	if db_init == "true" {
		init_data(DB)
	}
	return status
}

// func ConfigSession() *postgres.Storage {
// func ConfigSession() {

// 	host := config.Config("DB_HOST")
// 	port := config.Config("DB_PORT")
// 	user := config.Config("DB_USER")
// 	password := config.Config("DB_PASSWORD")
// 	name := config.Config("DB_NAME")
// 	sshmode := config.Config("SSH")
// 	post, _ := strconv.Atoi(port)

// 	// store := postgres.New(postgres.Config{
// 	Store = postgres.New(postgres.Config{

// 		Host:       host,
// 		Port:       post,
// 		Username:   user,
// 		Password:   password,
// 		Database:   name,
// 		Table:      "session",
// 		Reset:      false,
// 		GCInterval: 10 * time.Second,
// 		SslMode:    sshmode,
// 	})
// 	// return store
// }
