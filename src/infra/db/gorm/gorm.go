package gorm

import (
	"fmt"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {

	//Local development
	//connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Fortaleza"
	//DB, err = gorm.Open(postgres.Open(conectionString))

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require TimeZone=America/Fortaleza",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	productCategories := []models.ProductCategory{
		{Description: "Lanche"},
		{Description: "Acompanhamento"},
		{Description: "Bebida"},
		{Description: "Sobremesa"},
	}

	if !DB.Migrator().HasTable("product_categories") {
		DB.Migrator().CreateTable(&productCategories)
		DB.Create(&productCategories)
	}

	DB.AutoMigrate(&models.Product{})
}
