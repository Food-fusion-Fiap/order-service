package repositories

import (
	"context"
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	models "github.com/Food-fusion-Fiap/order-service/src/infra/db/models"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/mongo_driver"
	models2 "github.com/Food-fusion-Fiap/order-service/src/infra/db/mongo_driver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type OrderRepository struct {
}

func SetDefaultValues(sortBy string, orderBy string, status string) (string, string, string) {
	//TODO: sortBy, orderBy and status needs to be ENUMs, otherwise, it pops syntax error on log

	if sortBy == "" {
		sortBy = GetCreatedAtFieldName()
	}

	if orderBy == "" {
		orderBy = GetAscOrder()
	}

	return sortBy, orderBy, status
}

func (r OrderRepository) List(sortBy string, orderBy string, status string) ([]entities.Order, error) {
	var orderModels []models.Order

	sortBy, orderBy, status = SetDefaultValues(sortBy, orderBy, status)

	filter := bson.D{}
	if len(status) > 0 {
		filter = bson.D{{"status", status}}
	}

	sortOrder := 1
	if orderBy == GetDescOrder() {
		sortOrder = -1
	}
	opts := options.Find().SetSort(bson.D{{Key: sortBy, Value: sortOrder}})

	cursor, err := mongo_driver.OrdersCollection.Find(context.TODO(), filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &orderModels); err != nil {
		return nil, err
	}

	var orders []entities.Order
	for _, orderModel := range orderModels {
		orders = append(orders, *orderModel.ToDomain(orderModelProductsToProductInsideOrder(orderModel)))
	}
	return orders, nil
}

func (r OrderRepository) FindById(orderId string) (*entities.Order, error) {
	var orderModel models.Order

	objID, convErr := primitive.ObjectIDFromHex(orderId)

	if convErr != nil {
		return nil, errors.New("invalid orderId to convert to primitive.ObjectID")
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	err := mongo_driver.OrdersCollection.FindOne(context.TODO(), filter).Decode(&orderModel)

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return &entities.Order{}, nil
	} else if err != nil {
		return nil, err
	}

	productsOrderModel := orderModelProductsToProductInsideOrder(orderModel)
	return orderModel.ToDomain(productsOrderModel), nil
}

func orderModelProductsToProductInsideOrder(orderModel models.Order) []entities.ProductInsideOrder {
	var productsOrderModel []entities.ProductInsideOrder
	for _, p := range orderModel.Products {
		productsOrderModel = append(productsOrderModel, entities.ProductInsideOrder{
			Product:     p.Product.ToDomain(),
			Quantity:    p.Quantity,
			Observation: p.Observation,
		})
	}
	return productsOrderModel
}

func (r OrderRepository) Update(order *entities.Order) (*entities.Order, error) {

	objID, convErr := primitive.ObjectIDFromHex(order.ID)

	if convErr != nil {
		return nil, errors.New("invalid orderId to convert to primitive.ObjectID")
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: order.Status},
			{Key: "payment_status", Value: order.PaymentStatus},
			{Key: "updated_at", Value: primitive.NewDateTimeFromTime(time.Now().In(mongo_driver.LocaleApp))},
		}},
	}

	_, err := mongo_driver.OrdersCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return r.FindById(order.ID)
}

func (r OrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	var model models.Order

	model.Status = order.Status
	model.PaymentStatus = order.PaymentStatus

	var productsOrderModel []models.OrderProduct
	for _, p := range order.Products {
		productsOrderModel = append(productsOrderModel, models.OrderProduct{
			Product: models2.Product{
				ID:          int64(p.Product.ID),
				Name:        p.Product.Name,
				Price:       p.Product.Price,
				Description: p.Product.Description,
				CategoryID:  int64(p.Product.CategoryId),
			},
			Quantity:    p.Quantity,
			Observation: p.Observation,
		})
	}

	model.Products = productsOrderModel
	model.CreatedAt = primitive.NewDateTimeFromTime(time.Now().In(mongo_driver.LocaleApp))
	model.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().In(mongo_driver.LocaleApp))

	result, err := mongo_driver.OrdersCollection.InsertOne(context.TODO(), model)
	if err != nil {
		log.Println(err)
		return nil, errors.New("ocorreu um erro desconhecido ao criar o pedido")
	}

	model.ID = result.InsertedID.(primitive.ObjectID)
	return model.ToDomain(order.Products), nil
}

func GetDescOrder() string {
	return "DESC"
}

func (r OrderRepository) GetDescOrder() string {
	return GetDescOrder()
}

func GetAscOrder() string {
	return "ASC"
}

func (r OrderRepository) GetAscOrder() string {
	return GetAscOrder()
}

func GetCreatedAtFieldName() string {
	return "created_at"
}

func (r OrderRepository) GetCreatedAtFieldName() string {
	return GetCreatedAtFieldName()
}
