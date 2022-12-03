package database

import (
	"TerraInnAPI/config"
	"TerraInnAPI/model"
	"fmt"

	// postgresDriver "gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var Store *postgres.Storage

func Connect() bool {
	var err error
	status := true
	db_host := config.Config("DB_HOST")
	db_port := config.Config("DB_PORT")
	db_user := config.Config("DB_USER")
	db_password := config.Config("DB_PASSWORD")
	db_name := config.Config("DB_NAME")
	// db_ssh := config.Config("DB_SSH")
	// db_timezone := config.Config("APP_TIME_ZONE")

	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db_host, db_port, db_user, db_password, db_name, db_ssh)

	//MySQL DNS:
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("CONNECT DB ERROR: ", err)
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

	if config.Config("DB_REMIGRATE") == "true" {
		DB.Migrator().DropTable(&model.User{})
		// DB.Migrator().DropTable(&model.ReviewApplication{})
		// DB.Migrator().DropTable(&model.ApprovalFlow{})
		// DB.Migrator().DropTable(&model.ApprovalFlowDetail{})

		// Migrate the database
		DB.AutoMigrate(&model.User{})
	}

	if config.Config("DB_INIT") == "true" {
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
