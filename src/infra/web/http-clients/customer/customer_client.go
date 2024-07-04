package customer

import (
	"encoding/json"
	"fmt"
	"github.com/Food-fusion-Fiap/order-service/src/core/domain/entities"
	"io"
	"net/http"
)

type CustomerClient struct{}

func (c CustomerClient) Create(entity *entities.Customer) (*entities.Customer, error) {
	return entity, nil
}

func (c CustomerClient) List(entity *entities.Customer) ([]entities.Customer, error) {
	var customers []entities.Customer

	return customers, nil
}

func (c CustomerClient) FindFirstById(id uint) (*entities.Customer, error) {
	client := &http.Client{}

	resp, err := client.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	var customer entities.Customer

	err = json.Unmarshal(body, &customer)

	if err != nil {
		fmt.Printf("Error unmarshalling response body: %v\n", err)
		return nil, err
	}

	return &customer, nil
}
