package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Database struct
type Database struct {
	DB *gorm.DB
}

//NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	//USER := os.Getenv("DB_USER")
	//PASS := os.Getenv("DB_PASSWORD")
	//HOST := os.Getenv("DB_HOST")
	//PORT := os.Getenv("DB_PORT")
	//DBNAME := os.Getenv("DB_NAME")

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi", HOST, USER, PASS, DBNAME, PORT)

	dsn := "host=fullstack-postgres user=steven password=password dbname=fullstack_api port=5432 sslmode=disable TimeZone=Africa/Nairobi"

	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")

	}

	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
