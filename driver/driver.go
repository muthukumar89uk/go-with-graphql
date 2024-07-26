package driver

import (
	"fmt"
	dbmodel "jobPostGraphQl/db_model"
	"jobPostGraphQl/helper"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func DatabaseConnection() error {
	err := helper.Configure(".env")
	if err != nil {
		fmt.Println("error is loading env file")
		return err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	connectionURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	//connecting to postgres-SQL
	DBInstance, err = gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s,database connection sucessfull\n", dbname)
	return nil
}

// migrate and sync the model to create a db table
func MigrateDB() error {
	if err := DBInstance.AutoMigrate(&dbmodel.User{}, &dbmodel.Profile{}, &dbmodel.Career{}); err != nil {
		fmt.Println("failed to create the tables")
		return err
	}
	fmt.Println("table created successfully....")

	return nil
}
