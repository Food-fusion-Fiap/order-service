package repositories

import (
	"errors"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/gorm"
	"github.com/Food-fusion-Fiap/order-service/src/infra/db/models"
	gorm2 "gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
}

func (r ProductRepository) Create(entity *entities.Product) (*entities.Product, error) {
	product := models.Product{
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
		CategoryID:  entity.CategoryId,
	}

	dbResult := gorm.DB.Create(&product)

	if err := checkError(dbResult); err != nil {
		return nil, err
	}

	result := product.ToDomain()

	return &result, nil
}

func (r ProductRepository) FindById(id uint) (*entities.Product, error) {
	var product models.Product

	err := checkError(gorm.DB.Find(&product, id))

	if err != nil {
		return &entities.Product{}, err
	}

	result := product.ToDomain()

	return &result, nil
}

func (r ProductRepository) FindByIds(ids []uint) ([]entities.Product, error) {
	var products []models.Product
	var response []entities.Product
	gorm.DB.Where("id in (?)", ids).Find(&products)

	for _, product := range products {
		response = append(response, product.ToDomain())
	}

	return response, nil
}

func (r ProductRepository) FindAll() ([]entities.Product, error) {
	var products []models.Product

	err := checkError(gorm.DB.Find(&products))

	if err != nil {
		return []entities.Product{}, err
	}

	productEntities := []entities.Product{}

	for _, product := range products {
		productEntities = append(productEntities, product.ToDomain())
	}

	return productEntities, nil
}

func (r ProductRepository) FindByCategoryId(categoryId uint) ([]entities.Product, error) {
	var products []models.Product

	err := checkError(gorm.DB.Where(&models.Product{CategoryID: categoryId}).Find(&products))

	if err != nil {
		return []entities.Product{}, err
	}

	productEntities := []entities.Product{}

	for _, product := range products {
		productEntities = append(productEntities, product.ToDomain())
	}

	return productEntities, nil
}

func (p ProductRepository) DeleteById(id uint) error {
	var product models.Product

	err := checkError(gorm.DB.Delete(&product, id))

	if err != nil {
		return err
	}

	return nil
}

func (p ProductRepository) Edit(entity *entities.Product) (*entities.Product, error) {
	var product models.Product

	gorm.DB.Find(&product, entity.ID)

	product.PatchFields(entity.Name, entity.Price, entity.Description, entity.CategoryId)

	err := checkError(gorm.DB.Model(&product).Clauses(clause.Returning{}).UpdateColumns(&product))

	if err != nil {
		return nil, err
	}

	result := product.ToDomain()

	return &result, nil
}

func checkError(db *gorm2.DB) error {
	if err := db.Error; err != nil {
		message := "Houve um erro para realizar a persistência dos dados " + err.Error()
		return errors.New(message)
	}

	return nil
}
