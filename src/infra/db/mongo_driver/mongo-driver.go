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
)

var (
	DB                          *mongo.Client
	productCategoriesCollection *mongo.Collection
	productsCollection          *mongo.Collection
	ordersCollection            *mongo.Collection
	orderProductsCollection     *mongo.Collection
)

func ConnectDB() {

	// MongoDB connection string
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

	// Select the database and collection
	db := client.Database(os.Getenv("MONGO_INITDB_DATABASE"))
	productCategoriesCollection = db.Collection("product_categories")
	productsCollection = db.Collection("products")
	ordersCollection = db.Collection("orders")
	orderProductsCollection = db.Collection("order_products")

	// Insert default product categories if they don't exist
	productCategories := []models.ProductCategory{
		{Description: "Lanche"},
		{Description: "Acompanhamento"},
		{Description: "Bebida"},
		{Description: "Sobremesa"},
	}

	for _, category := range productCategories {
		filter := bson.D{{Key: "description", Value: category.Description}}
		count, err := productCategoriesCollection.CountDocuments(context.TODO(), filter)
		if err != nil {
			log.Println("Erro ao verificar existência de categoria:", err)
			continue
		}
		if count == 0 {
			_, err := productCategoriesCollection.InsertOne(context.TODO(), category)
			if err != nil {
				log.Println("Erro ao inserir categoria:", err)
			}
		}
	}
}
