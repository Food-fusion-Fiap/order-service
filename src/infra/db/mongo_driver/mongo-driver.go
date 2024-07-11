package mongo_driver

import (
	"context"
	"fmt"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	ProductCategoriesCollection *mongo.Collection
	OrdersCollection            *mongo.Collection
	LocaleApp                   *time.Location
)

func ConnectDB() {
	locale, err := time.LoadLocation("America/Sao_Paulo")
	LocaleApp = locale

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017",
		os.Getenv("MONGO_INITDB_ROOT_USERNAME"), os.Getenv("MONGO_INITDB_ROOT_PASSWORD"), os.Getenv("MONGO_INITDB_HOST"))

	fmt.Println(uri)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados MongoDB:", err)
	}

	fmt.Println("Conectado ao MongoDB!")

	db := client.Database(os.Getenv("MONGO_INITDB_DATABASE"))
	ProductCategoriesCollection = db.Collection("product_categories")
	OrdersCollection = db.Collection("orders")

	productCategories := []models.ProductCategory{
		{Description: "Lanche"},
		{Description: "Acompanhamento"},
		{Description: "Bebida"},
		{Description: "Sobremesa"},
	}

	for _, category := range productCategories {
		filter := bson.D{{Key: "description", Value: category.Description}}
		count, err := ProductCategoriesCollection.CountDocuments(context.TODO(), filter)
		if err != nil {
			log.Println("Erro ao verificar existência de categoria:", err)
			continue
		}
		if count == 0 {
			_, err := ProductCategoriesCollection.InsertOne(context.TODO(), category)
			if err != nil {
				log.Println("Erro ao inserir categoria:", err)
			}
		}
	}
}
