package config

import (
	"fmt"
	"os"

	"github.com/Damir564/HttpApiGo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Printf("host=localhost user=%s password=%s dbname=%s port=5432", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	// dsn := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	err = db.SetupJoinTable(&models.User{}, "Segments", &models.UserSegments{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	// s1 := db.Model(&models.UserSegments{}).Association("SegmentId").Relationship.
	// s2 := db.Model(&models.UserSegments{}).Association("SegmentId").Relationship.ParseConstraint().OnDelete
	// fmt.Println(s1, " ", s2)
	DB = db

}
